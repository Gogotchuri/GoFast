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
	apiGroup := app.Group("/api") //TODO:ratelimiters
	//TODO: initialize other api routes here (Decomposition recommended)
	//Last api route maps all unmapped routes to one handler
	apiGroup.Get("/*", func(c *fiber.Ctx) {
		c.Send("First api route (index).")
	})
	apiGroup.Post("/*", func(c *fiber.Ctx) {
		c.Send("First api route (index).")
	})
}

func initStaticRoutes(app *fiber.App) {
	app.Static("/", "./public/")
	app.Get("/*", func(ctx *fiber.Ctx) {
		ctx.SendFile("./public/index.html")
	})
}
