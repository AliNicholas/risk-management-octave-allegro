package main

import (
	"final_allegro/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./src")

	app.Get("/", handlers.Home)
	app.Get("/form", handlers.Form)
	// app.Post("/form")
	// app.Post("/delete", func(c *fiber.Ctx) error {
	// 	return c.Render("index", fiber.Map{
	// 		"Value": "Hello",
	// 	}, "delete-btn")
	// })

	log.Fatal(app.Listen(":3000"))
}
