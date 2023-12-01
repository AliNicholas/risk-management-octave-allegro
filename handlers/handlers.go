package handlers

import (
	"final_allegro/models"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// LOGIN PAGE
func LoginPage(ctx *fiber.Ctx) error {
	return ctx.Render("login", fiber.Map{})
}

func Logout(ctx *fiber.Ctx) error {
	return ctx.Redirect("/")
}
func Login(ctx *fiber.Ctx) error {
	return ctx.SendString("ok")
}
func Register(ctx *fiber.Ctx) error {
	return ctx.SendString("ok")
}

// Dashboard Page
func DashboardPage(ctx *fiber.Ctx) error {
	var Projects = []models.Project{
		{Title: "Project 1", ID: 1},
		{Title: "Project 10", ID: 2},
		// ... tambahkan proyek lainnya sesuai kebutuhan
	}

	var Prior = models.ImpactPriority{
		ReputationConfidence: 1,
		Financial:            2,
		Productivity:         3,
		SafetyHealth:         4,
		FinesLegalPenalties:  5,
	}

	return ctx.Render("dashboard", fiber.Map{
		"TotalAssets": 1,
		"TotalRisks":  1,
		"Projects":    Projects,
		"Prior":       Prior,
	})
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

// Form Page
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

// Asset Container Page
func AddProfilePage(ctx *fiber.Ctx) error {
	return ctx.Render("add-asset-container", nil)
}
func PostProfile(ctx *fiber.Ctx) error {
	return ctx.SendString("ok")
}

// Risk Page
func AddRiskPage(ctx *fiber.Ctx) error {
	return ctx.Render("add-risk", nil)
}
func AddRiskPost(ctx *fiber.Ctx) error {
	return ctx.SendString("ok")
}

func HandleScore(c *fiber.Ctx) error {
	projectIdStr := c.Params("projectId")
	area := c.Params("area")
	value := c.FormValue("area")

	projectId, err := strconv.ParseUint(projectIdStr, 10, 64)
	if err != nil {
		panic(err)
	}

	result, err := models.GetPriorityByProjectId(uint(projectId))
	if err != nil {
		panic(err)
	}

	area1 := result.ReputationConfidence
	area2 := result.Financial
	area3 := result.Productivity
	area4 := result.SafetyHealth
	area5 := result.FinesLegalPenalties

	var x int

	if area == "area1" {
		if value == "High" {
			x = area1 * 3
		} else if value == "Medium" {
			x = area1 * 2
		} else if value == "Low" {
			x = area1 * 1
		}
	} else if area == "area2" {
		if value == "High" {
			x = area2 * 3
		} else if value == "Medium" {
			x = area2 * 2
		} else if value == "Low" {
			x = area2 * 1
		}
	} else if area == "area3" {
		if value == "High" {
			x = area3 * 3
		} else if value == "Medium" {
			x = area3 * 2
		} else if value == "Low" {
			x = area3 * 1
		}
	} else if area == "area4" {
		if value == "High" {
			x = area4 * 3
		} else if value == "Medium" {
			x = area4 * 2
		} else if value == "Low" {
			x = area4 * 1
		}
	} else if area == "area5" {
		if value == "High" {
			x = area5 * 3
		} else if value == "Medium" {
			x = area5 * 2
		} else if value == "Low" {
			x = area5 * 1
		}
	}

	response := fmt.Sprintf("<p>%s</p>", strconv.Itoa(x))
	return c.SendString(response)
}
