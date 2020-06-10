package auth

import (
	"net/http"

	"github.com/Gogotchuri/GoFast/app/models"
	"github.com/Gogotchuri/GoFast/app/services/validators"

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
	if models.GetUserByEmail(req.Email) != nil {
		c.Status(http.StatusUnauthorized).JSON("A user with the entered Email already exists!")
		return
	}

	user := models.User{
		Email:     req.Email,
		Password:  req.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}
	user.Save()

	c.Status(http.StatusCreated).JSON(user)
}
