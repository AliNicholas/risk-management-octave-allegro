package main

import (
	"final_allegro/database"
	"final_allegro/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	database.ConnectDB()

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./src")

	app.Get("/", handlers.LoginPage)
	app.Get("/dashboard", handlers.DashboardPage)

	app.Get("/form", handlers.FormPage)
	app.Post("/form", handlers.FormPost)
	app.Get("/notif", handlers.Notification)

	app.Get("/add-asset", handlers.AddAssetPage)
	app.Post("/add-asset", handlers.AddAssetPost)

	app.Get("/add-container", handlers.AddContainerPage)
	app.Post("/add-container", handlers.AddContainerPost)

	app.Get("/add-risk", handlers.AddRiskPage)
	app.Post("/add-risk", handlers.AddRiskPost)

	app.Delete("/delete-project", handlers.DeleteProject)

	// tba
	app.Get("/logout", handlers.Logout)
	app.Post("/login", handlers.Login)
	app.Post("/register", handlers.Register)

	log.Fatal(app.Listen(":3000"))

}
