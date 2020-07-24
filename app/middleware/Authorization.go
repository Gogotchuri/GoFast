package middleware

import (
	"github.com/Gogotchuri/GoFast/app/models"
	"github.com/Gogotchuri/GoFast/app/services"
	"github.com/Gogotchuri/GoFast/app/services/errors"

	"github.com/gofiber/fiber"
)

/*IsAuthorized Check is user is authorized*/
func IsAuthorized() func(c *fiber.Ctx) {
	return func(c *fiber.Ctx) {
		accessDetails, err := services.JWTHasValidToken(services.JWTExtractToken(c.Get("Authorization")))
		if err != nil || accessDetails == nil{
			errors.SendUnauthorized(c)
			return
		}
		user := models.GetUserByID(accessDetails.UserID)
		if user == nil {
			errors.SendUnauthorized(c)
			return
		}
		c.Locals("jwt", accessDetails)
		c.Locals("user", user) //Set user as a local
		c.Next()
	}
}
