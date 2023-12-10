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

	app.Get("/contoh", func(c *fiber.Ctx) error {
		return c.Render("contoh", nil)
	})

	// Login
	// app.Get("/", handlers.LoginPage)
	// app.Post("/register", handlers.Register)
	// app.Post("/login", handlers.Login)
	// app.Get("/logout", handlers.Logout)

	// Dashboard
	app.Get("//:selector?", handlers.DashboardPage)
	app.Get("/render", handlers.Render)
	app.Delete("/project/:projectId", handlers.DeleteProject)
	app.Get("/asset/:projectId", handlers.RenderByAsset)
	app.Delete("/asset/:assetId", handlers.DeleteAsset)

	// Form
	app.Get("/form", handlers.FormPage)
	app.Post("/form", handlers.FormPost)
	app.Get("/notif", handlers.Notification)

	// Asset Profile
	app.Get("/add-profile/:projectId", handlers.AddProfilePage)
	app.Post("/add-profile/:projectId", handlers.PostProfile)

	// Risk
	app.Get("/add-risk/:projectId", handlers.AddRiskPage)
	app.Post("/asset", handlers.RenderAsset)
	app.Post("/add-risk/:projectId", handlers.PostRisk)
	app.Post("/score/:projectId/:area/:recentScore", handlers.HandleScore)

	// end

	log.Fatal(app.Listen(":3000"))
}
