package auth

import (
	"github.com/Gogotchuri/GoFast/app/services"
	"github.com/Gogotchuri/GoFast/app/services/errors"
	"net/http"
	"time"

	"github.com/Gogotchuri/GoFast/app/models"
	"github.com/Gogotchuri/GoFast/app/services/cache"
	"github.com/Gogotchuri/GoFast/app/services/hash"
	"github.com/Gogotchuri/GoFast/app/services/misc"
	"github.com/Gogotchuri/GoFast/app/services/validators"
	"github.com/Gogotchuri/GoFast/config"

	"github.com/gofiber/fiber"
)

/*SignIn Signs user into their account*/
func SignIn(c *fiber.Ctx) {
	var req validators.SignInRequestT
	// Parse input
	if err := c.BodyParser(&req); err != nil {
		errors.SendDefaultUnprocessable(c)
		return
	}

	// Check for invalid input
	if errs := req.Validate(); errs != nil {
		errors.SendErrors(c, http.StatusUnprocessableEntity, errs)
		return
	}
	// Get user with email-password combination
	user := models.CheckCredentials(req.Email, req.Password)
	if user == nil {
		errors.SendErrors(c, http.StatusUnauthorized, &[]string{"Invalid credentials"})
		return
	}
	tokensJSON := CreateTokensForUser(c, user)
	if tokensJSON == nil {
		return
	}
	c.Status(http.StatusAccepted).JSON(*tokensJSON)
}

/*SignUp Creates a new account, if credentials are valid*/
func SignUp(c *fiber.Ctx) {
	otacConf := config.GetInstance().Redis.OTAC

	var req validators.SignUpRequestT
	// Parse input
	if err := c.BodyParser(&req); err != nil {
		errors.SendDefaultUnprocessable(c)
		return
	}

	// Check for invalid input
	if errs := req.Validate(); errs != nil {
		errors.SendErrors(c, http.StatusUnprocessableEntity, errs)
		return
	}
	// Check if a user with passed mail already exists
	if errs := models.GetUserByEmail(req.Email); errs != nil {
		errors.SendErrors(c, http.StatusUnprocessableEntity, &[]string{"Provided email is already taken"})
		return
	}

	cachedOTAC, err := cache.GetRedisInstance().Get(otacConf.EntryPrefix + req.Email).Result()
	if err != nil || cachedOTAC != req.OTAC {
		c.Status(http.StatusUnauthorized).JSON("Entered verification code was wrong or it has expired!")
		return
	}

	user := models.User{
		Email:     req.Email,
		Password:  hash.GetPasswordHash(req.Password),
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}
	user.Save()

	_ = c.Status(http.StatusCreated).JSON(user.ToMap())
}

/*SendVerificationMail Generates random code and sends to passed email*/
func SendVerificationMail(c *fiber.Ctx) {
	otacConf := config.GetInstance().Redis.OTAC

	var req validators.VerificationRequestT
	// Parse input
	if err := c.BodyParser(&req); err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON("Request parsing failed!")
		return
	}

	// Check for invalid input
	if errs := req.Validate(); errs != nil {
		c.Status(http.StatusUnauthorized).JSON(fiber.Map{"errors": *errs})
		return
	}
	// Check if a user with passed mail already exists
	if errs := models.GetUserByEmail(req.Email); errs != nil {
		c.Status(http.StatusUnauthorized).JSON("A user with the entered Email already exists!")
		return
	}

	code := misc.RandCode()

	if err := cache.GetRedisInstance().Set(otacConf.EntryPrefix+req.Email, code, time.Duration(otacConf.Expires)*time.Second).Err(); err != nil {
		c.Status(http.StatusInternalServerError).JSON(err)
	}

	go misc.SendMail(req.Email, "Your verification code is: "+ code, "GoFast email verification")

	c.Status(http.StatusCreated).JSON(code)
}

func GetUserDetails(c *fiber.Ctx) {
	user := c.Locals("user").(*models.User)
	c.JSON(user.ToMap())
}

/*Returns map of new jwt access, refresh tokens and user json for given user*/
func CreateTokensForUser(c *fiber.Ctx, user *models.User) *map[string]interface{} {
	tokens, err := services.JWTCreateToken(user.ID)
	if err != nil {
		errors.SendErrors(c, http.StatusUnprocessableEntity, &[]string{err.Error()})
		return nil
	}

	if sErr := tokens.Save(); sErr != nil {
		errors.SendErrors(c, http.StatusUnprocessableEntity, &[]string{sErr.Error()})
		return nil
	}

	tokensJSON := map[string]interface{}{
		"access_token":  tokens.Access.Token,
		"refresh_token": tokens.Refresh.Token,
		"user": user.ToMap(),
	}

	return &tokensJSON
}