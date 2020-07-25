package auth

import (
	"github.com/Gogotchuri/GoFast/app/models"
	"github.com/Gogotchuri/GoFast/app/services/errors"
	"github.com/Gogotchuri/GoFast/config"

	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gofiber/fiber"
	"gopkg.in/danilopolani/gocialite.v1"
)

var gocial *gocialite.Dispatcher = nil
var onceGocial sync.Once

func getGocialInstance() *gocialite.Dispatcher {
	onceGocial.Do(func() {
		gocial = gocialite.NewDispatcher()
	})
	return gocial
}

func CallbackHandler(c *fiber.Ctx) {
	c.Set("content-type", "text/html") //Setting correct content type to render passed script

	provider := c.Params("provider")
	if provider != "google" && provider != "facebook" {
		c.Send(jsPostToWindowOpener("", "", "Unsupported provider"))
	}
	// Retrieve query params for state and code
	state := c.Query("state")
	code := c.Query("code")
	// Handle callback and check for errors
	socialUser, _, err := getGocialInstance().Handle(state, code)
	if err != nil {
		c.Send(jsPostToWindowOpener("", "", "Couldn't process social sign in request"))
		return
	}

	user := models.GetUserByEmail(socialUser.Email)
	if user == nil {
		user = &models.User{
			Email:     socialUser.Email,
			FirstName: socialUser.FirstName,
			LastName:  socialUser.LastName,
			EmailVerifiedAt: time.Now(),
		}

		if provider == "facebook" {
			user.FacebookID = socialUser.ID
		} else if provider == "google" {
			user.GoogleID = socialUser.ID
		} else {
			c.Send(jsPostToWindowOpener("", "", "Unsupported provider"))
			return
		}
		user.Save()
	} else {
		//enforce uniqueness of ids
		if provider == "facebook" {
			if user.FacebookID != socialUser.ID {
				c.Send(jsPostToWindowOpener("", "", "Your email is already used by other user"))
				return
			}
		} else if provider == "google" {
			if user.GoogleID != socialUser.ID {
				c.Send(jsPostToWindowOpener("", "", "Your email is already used by other user"))
				return
			}
		} else {
			//Sending js script with embedded empty tokens and error message to trigger websocket event on window opener
			c.Send(jsPostToWindowOpener("", "", "Unsupported provider"))
			return
		}
	}

	tokensJSON := CreateTokensForUser(c, user)
	if tokensJSON == nil {
		return
	}

	//Sending js script with embedded tokens to trigger websocket event on window opener
	c.Send(jsPostToWindowOpener((*tokensJSON)["access_token"].(string), (*tokensJSON)["refresh_token"].(string), ""))
}

func jsPostToWindowOpener(accessToken, refreshToken, error string) string{
	return fmt.Sprintf("<script>\n" +
		"window.opener.postMessage({ access_token: '%s', refresh_token: '%s', error:'%s' }, '*');\n" + //TODO:set origins strictly
		"window.close();\n " +
		"</script>", accessToken, refreshToken, error)
}

/*Returns url to be redirected to for social authentication*/
func RedirectHandler(c *fiber.Ctx) {
	provider := c.Params("provider")
	if provider != "google" && provider != "facebook" {
		errors.SendErrors(c, http.StatusUnprocessableEntity, &[]string{"unsupported provider"})
	}
	cfg := config.GetInstance()
	providerSecrets := map[string]map[string]string{
		"facebook": {
			"clientID":     cfg.Facebook.ClientID,
			"clientSecret": cfg.Facebook.ClientSecret,
			"redirectURL":  cfg.Domain+"/api/auth/facebook/callback",
		},
		"google": {
			"clientID":    cfg.Google.ClientID,
			"clientSecret": cfg.Google.ClientSecret,
			"redirectURL":  cfg.Domain + "/api/auth/google/callback",
		},
	}
	providerData := providerSecrets[provider]
	authURL, err := getGocialInstance().New().
		Driver(provider).
		Redirect(
			providerData["clientID"],
			providerData["clientSecret"],
			providerData["redirectURL"],
		)

	// Check for errors (usually driver not valid)
	if err != nil {
		errors.SendErrors(c, http.StatusInternalServerError, &[]string{err.Error()})
		return
	}
	// Redirect with authURL
	c.Status(http.StatusOK).JSON(map[string]string{"url" : authURL})
}
