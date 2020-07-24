package errors

import (
	"net/http"

	"github.com/gofiber/fiber"
)

/*SendDefaultUnprocessable sends 422 error code to the client with default error message*/
func SendDefaultUnprocessable(ctx *fiber.Ctx) {
	SendErrors(ctx, http.StatusUnprocessableEntity, &[]string{"Couldn't parse request"})
}

/*SendUnauthorized sends 401 error code to the client*/
func SendUnauthorized(ctx *fiber.Ctx) {
	SendErrors(ctx, http.StatusUnauthorized, &[]string{"Unauthorized"})
}

/*SendErrors sends error with status to the client with errors array*/
func SendErrors(ctx *fiber.Ctx, status int, errors *[]string) {
	ctx.Status(status).JSON(fiber.Map{"errors": *errors})
}
