package main

import (
	"github.com/Gogotchuri/GoFast/config"
	"github.com/Gogotchuri/GoFast/database"
	"github.com/Gogotchuri/GoFast/database/migrations"
	"github.com/Gogotchuri/GoFast/routes"

	"github.com/gofiber/fiber"
)

func main() {
	//Setup database
	db := database.GetInstance()
	defer db.Close() //Close database
	err := migrations.DefaultMigration(db)
	mustPanic(err)
	//Create fiber app
	app := fiber.New()
	//Setup routes
	routes.InitializeRoutes(app)
	err = app.Listen(config.GetInstance().Port)
	mustPanic(err)
}

func mustPanic(err error) {
	if err != nil {
		panic(err)
	}
}
