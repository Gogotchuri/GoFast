package main

import (
	"github.com/Gogotchuri/GoFast/routes"
	"github.com/gofiber/fiber"
)

func main(){
	app := fiber.New()
	routes.InitializeRoutes(app)
	err := app.Listen(8081)
	mustPanic(err)
}

func mustPanic(err error) {
	if err != nil {
		panic(err)
	}
}