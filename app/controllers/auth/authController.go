package auth

import (
	"net/http"

	"github.com/Gogotchuri/GoFast/app/services"
	"github.com/Gogotchuri/GoFast/app/services/errors"

	"github.com/Gogotchuri/GoFast/app/events"
	"github.com/Gogotchuri/GoFast/app/models"
	"github.com/Gogotchuri/GoFast/app/services/hash"
	"github.com/Gogotchuri/GoFast/app/services/validators"

	"github.com/gofiber/fiber"
)

/*PasswordForgotten Sends verification link to passed email for resetting password*/
func PasswordForgotten(c *fiber.Ctx) {
	email := struct {
		Email string `json:"email"`
	}{}
	if err := c.BodyParser(&email); err != nil {
		errors.SendDefaultUnprocessable(c)
		return
	}

	user := models.GetUserByEmail(email.Email)
	if user == nil {
		_ = c.JSON("") //To prevent indexing we don't return error
		return
	}
	services.SendPasswordResetMail(user)
	_ = c.JSON("")
}

/*ResetPassword Updates user's password*/
func ResetPassword(c *fiber.Ctx) {
	rr := struct {
		Password string `json:"password"`
		Token    string `json:"token"`
	}{}
	if err := c.BodyParser(&rr); err != nil {
		errors.SendDefaultUnprocessable(c)
		return
	}
	if rr.Password == "" || rr.Token == "" {
		errors.SendErrors(c, http.StatusUnauthorized, &[]string{"Invalid credentials"})
		return
	}
	if !services.SetNewPassword(rr.Token, rr.Password) {
		errors.SendErrors(c, http.StatusUnauthorized, &[]string{"Reset link has expired"})
		return
	}
	_ = c.JSON("Password reset ended successfully")
}

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

	user := models.User{
		Email:     req.Email,
		Password:  hash.GetPasswordHash(req.Password),
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}
	user.Save()
	events.FireUserCreated(&user)
	_ = c.Status(http.StatusCreated).JSON(user.ToMap())
}

/*RefreshJWTTokens refreshes access token from existing refresh token*/
func RefreshJWTTokens(c *fiber.Ctx) {
	fieldMap := map[string]string{}
	if err := c.BodyParser(&fieldMap); err != nil {
		errors.SendDefaultUnprocessable(c)
		return
	}
	refreshToken := fieldMap["refresh_token"]
	tokDetails, err := services.JWTHasValidRefreshToken(refreshToken)
	if err != nil {
		errors.SendErrors(c, http.StatusUnauthorized, &[]string{err.Error()})
		return
	}

	err = tokDetails.Delete()
	if err != nil {
		errors.SendErrors(c, http.StatusUnauthorized, &[]string{"Unauthorized", "Refresh Token Expired"})
		return
	}
	user := models.GetUserByID(tokDetails.UserID)
	if user == nil {
		errors.SendErrors(c, http.StatusUnauthorized, &[]string{"Unauthorized", "Not a valid user"})
		return
	}
	tokensJSON := CreateTokensForUser(c, user)
	if tokensJSON == nil {
		return
	}
	_ = c.Status(http.StatusAccepted).JSON(*tokensJSON)
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
		"user":          user.ToMap(),
	}

	return &tokensJSON
}

/*Logout logs out user from system*/
func Logout(c *fiber.Ctx) {
	//Extract token details
	var accessTD = c.Locals("jwt").(*services.JWTAccessDetails)
	if delErr := accessTD.Delete(); delErr != nil {
		errors.SendUnauthorized(c)
		return
	}
	c.Status(http.StatusOK).JSON("Successfully logged out")
}
