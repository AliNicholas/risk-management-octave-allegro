package handlers

import "github.com/gofiber/fiber/v2"

func Home(ctx *fiber.Ctx) error {
	return ctx.Render("index", nil)
}

func Form(ctx *fiber.Ctx) error {
	return ctx.Render("form", nil)
}
