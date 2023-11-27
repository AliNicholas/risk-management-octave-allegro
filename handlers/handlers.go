package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func Home(ctx *fiber.Ctx) error {
	return ctx.Render("index", nil)
}

func Form(ctx *fiber.Ctx) error {
	return ctx.Render("form", nil)
}

func Risk(ctx *fiber.Ctx) error {
	return ctx.Render("risk", nil)
}

func Detail(ctx *fiber.Ctx) error {
	return ctx.Render("detail", nil)
}

func AddNewProject(ctx *fiber.Ctx) error {
	return ctx.SendString("ok")
}

func AddNewRisk(ctx *fiber.Ctx) error {
	return ctx.SendString("ok")
}

func SelectProject(ctx *fiber.Ctx) error {
	return ctx.SendString("ok")
}

func SelectCategory(ctx *fiber.Ctx) error {
	return ctx.SendString("ok")
}
