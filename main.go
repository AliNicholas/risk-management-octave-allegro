package main

import (
	"errors"
	"final_allegro/handlers"
	"fmt"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type Counter struct {
	value int
	sync.Mutex
}

func (c *Counter) Increase() {
	c.Lock()
	c.value++
	c.Unlock()
}

func (c *Counter) Decrease() {
	c.Lock()
	c.value--
	c.Unlock()
}

func (c *Counter) GetValue() int {
	c.Lock()
	defer c.Unlock()
	return c.value
}

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
		// Override default error handler
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			// Send custom error page
			err = ctx.Status(code).SendFile(fmt.Sprintf("./%d.html", code))
			if err != nil {
				// In case the SendFile fails
				return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}

			// Return from handler
			return nil
		},
	})

	app.Static("/", "./src")

	app.Get("/", handlers.Home)
	app.Get("/form", handlers.Form)
	//app.Get("/", handlers.Home)

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}

//app.Post("/increase", func(ctx *fiber.Ctx) error {
//	tmplStr := "<div id=\"counter\">{{.CounterValue}}</div>"
//	tpl := template.Must(template.New("counter").Parse(tmplStr))
//	counter.Increase()
//	return tpl.ExecuteTemplate(ctx.Response().BodyWriter(), "counter", fiber.Map{
//		"CounterValue": counter.GetValue(),
//	})
//})
//
//app.Post("/increase", func(ctx *fiber.Ctx) error {
//	tmplStr := "<div id=\"counter\">{{.CounterValue}}</div>"
//	tpl := template.Must(template.New("counter").Parse(tmplStr))
//	counter.Increase()
//	return tpl.ExecuteTemplate(ctx.Response().BodyWriter(), "counter", fiber.Map{
//		"CounterValue": counter.GetValue(),
//	})
//
//})
//app.Post("/decrease", func(ctx *fiber.Ctx) error {
//	tmplStr := "<div id=\"counter\">{{.CounterValue}}</div>"
//	tpl := template.Must(template.New("counter").Parse(tmplStr))
//	counter.Decrease()
//	return tpl.ExecuteTemplate(ctx.Response().BodyWriter(), "counter", fiber.Map{
//		"CounterValue": counter.GetValue(),
//	})
//})
