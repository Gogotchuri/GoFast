package auth

import (
	"net/http"

	"github.com/Gogotchuri/GoFast/app/services/validators"

	"github.com/Gogotchuri/GoFast/app/models"
	"github.com/Gogotchuri/GoFast/app/services"
	"github.com/gofiber/fiber"
)

/*VerifyEmail Verifies user's mail by comparing received OTAC code to the on in cache*/
func VerifyEmail(c *fiber.Ctx) {
	user := c.Locals("user").(*models.User)

	var req validators.VerificationRequestT

	// Parse input
	if err := c.BodyParser(&req); err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON("Request parsing failed!")
		return
	}

	if services.VerifyEmail(user, req.OTAC) {
		// TODO: Wrong status?
		c.Status(http.StatusAccepted).JSON("Verification successful")
	} else {
		c.Status(http.StatusExpectationFailed).JSON("Entered verification code was wrong or it has expired!")
	}
}

/*SendVerificationMail Generates random code and sends to passed email*/
func ReSendVerificationMail(c *fiber.Ctx) {
	user := c.Locals("user").(*models.User)
	services.SendVerificationMail(user)
	c.Status(http.StatusCreated).JSON("Email verification sent")
}
