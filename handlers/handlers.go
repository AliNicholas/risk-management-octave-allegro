package handlers

import (
	"final_allegro/models"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func LoginPage(ctx *fiber.Ctx) error {
	return ctx.Render("login", fiber.Map{})
}

func DashboardPage(ctx *fiber.Ctx) error {
	var ProjectList []string

	return ctx.Render("dashboard", fiber.Map{
		// "TotalAssets": 1,
		"TotalRisks":  0,
		"ProjectList": ProjectList,
	})
}

func FormPage(ctx *fiber.Ctx) error {
	return ctx.Render("form", nil)
}

func FormPost(ctx *fiber.Ctx) error {
	// body := ctx.FormValue("ok")
	return ctx.SendString(`<div class="overlay" id="overlay"><div class="notification-box"><h4>Success</h4><a href="/notif">ok!</a></div></div>`)
}

func Notification(ctx *fiber.Ctx) error {
	return ctx.Redirect("dashboard")
}

func AddAssetPage(ctx *fiber.Ctx) error {
	return ctx.Render("add-asset", nil)
}
func AddAssetPost(ctx *fiber.Ctx) error {
	return ctx.SendString("ok")
}

func AddContainerPage(ctx *fiber.Ctx) error {
	return ctx.Render("add-container", nil)
}
func AddContainerPost(ctx *fiber.Ctx) error {
	return ctx.SendString("ok")
}

func AddRiskPage(ctx *fiber.Ctx) error {
	return ctx.Render("add-risk", nil)
}
func AddRiskPost(ctx *fiber.Ctx) error {
	return ctx.SendString("ok")
}

func DeleteProject(ctx *fiber.Ctx) error {
	strId := ctx.Get("id")

	id, err := strconv.ParseUint(strId, 10, 0)
	if err != nil {
		fmt.Println("Error parsing string to uint:", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid project ID",
		})
	}

	err = models.DeleteProjectById(uint(id))

	if err != nil {
		fmt.Println("Error deleting project:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete project",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Project deleted successfully",
	})
}

// tba
func Logout(ctx *fiber.Ctx) error {
	return ctx.Redirect("/")
}
func Login(ctx *fiber.Ctx) error {
	return ctx.SendString("ok")
}
func Register(ctx *fiber.Ctx) error {
	return ctx.SendString("ok")
}
