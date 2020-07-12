package auth

import (
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
		c.Status(http.StatusUnprocessableEntity).JSON("Request parsing failed!")
		return
	}

	// Check for invalid input
	if errs := req.Validate(); errs != nil {
		c.Status(http.StatusUnauthorized).JSON(fiber.Map{"errors": *errs})
		return
	}
	// Get user with email-password combination
	user := models.CheckCredentials(req.Email, req.Password)
	if user == nil {
		c.Status(http.StatusUnauthorized).JSON("Entered Email or Password is incorrect!")
		return
	}

	c.Status(http.StatusOK).JSON("Successfully logged in!")
}

/*SignUp Creates a new account, if credentials are valid*/
func SignUp(c *fiber.Ctx) {
	otacConf := config.GetInstance().Redis.OTAC

	var req validators.SignUpRequestT
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

	c.Status(http.StatusCreated).JSON(user)
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
		// TODO: What status should we use here?
		c.Status(http.StatusInternalServerError).JSON(err)
	}

	if err := misc.SendMail(req.Email, "Your verification code is: "+code, "GoFast email verfication"); err != "" {
		// TODO: What status should we use here?
		c.Status(http.StatusInternalServerError).JSON(err)
	}
	c.Status(http.StatusCreated).JSON(code)
}
