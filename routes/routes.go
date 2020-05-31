package routes

import (
	"github.com/gofiber/fiber"
)

/*InitializeRoutes given a fiber.App initializes routes on it*/
func InitializeRoutes(app *fiber.App) {
	initAPIRoutes(app)
	initStaticRoutes(app)
}

func initAPIRoutes(app *fiber.App) {
	apiGroup := app.Group("/api")
	//TODO: map other routes here

	//Last api routes maps all unmapped routes to one handler
	apiGroup.Post("/*", func(c *fiber.Ctx) {
		c.Send("Base api route for POST.")
	})
	apiGroup.Get("/*", func(c *fiber.Ctx) {
		c.Send("Base api route for GET.")
	})
}

func initStaticRoutes(app *fiber.App) {
	app.Static("/", "./public/") //Serving static content from here
	app.Get("/*", func(ctx *fiber.Ctx) { //Serving base html page from here
		ctx.SendFile("./public/index.html")
	})
}
