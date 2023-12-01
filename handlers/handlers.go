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
func DashboardPage(c *fiber.Ctx) error {

	projectIdStr := c.FormValue("selector")

	projects, err := models.GetAllProjects()
	if err != nil {
		return err
	}

	var totalAssets int
	var totalRisks int
	var category1 int
	var category2 int
	var category3 int
	var category4 int
	var selected int

	var assets []models.AssetInformation
	var risks models.AssetRisk
	var areaPriority models.ImpactPriority

	if projectIdStr != "" {
		projectId, err := strconv.Atoi(projectIdStr)
		if err != nil {
			panic(err)
		}

		selected = projectId

		assets, err = models.GetAllAssetByProjectId(uint(projectId))
		if err != nil {
			panic(err)
		}

		totalAssets = len(assets)

		risks, err := models.GetAllRisksByProjectId(uint(projectId))
		if err != nil {
			panic(err)
		}

		totalRisks = len(risks)

		areaPriority, err = models.GetPriorityByProjectId(uint(projectId))
		if err != nil {
			panic(err)
		}
	}
	return c.Render("dashboard", fiber.Map{
		"TotalAssets":  totalAssets,
		"TotalRisks":   totalRisks,
		"Projects":     projects,
		"AreaPriority": areaPriority,
		"ProjectId":    selected,
		"Category1":    category1,
		"Category2":    category2,
		"Category3":    category3,
		"Category4":    category4,
		"Risks":        risks,
		"Assets":       assets,
		"Selected":     selected,
	})
}

func Render(c *fiber.Ctx) error {
	projectIdStr := c.Params("project-selector")

	projects, err := models.GetAllProjects()
	if err != nil {
		return err
	}

	var totalAssets int
	var totalRisks int
	var category1 int
	var category2 int
	var category3 int
	var category4 int
	var selected int

	var assets []models.AssetInformation
	var risks models.AssetRisk
	var areaPriority models.ImpactPriority

	if projectIdStr != "" {
		projectId, err := strconv.Atoi(projectIdStr)
		if err != nil {
			panic(err)
		}

		selected = projectId

		assets, err = models.GetAllAssetByProjectId(uint(projectId))
		if err != nil {
			panic(err)
		}

		totalAssets = len(assets)

		risks, err := models.GetAllRisksByProjectId(uint(projectId))
		if err != nil {
			panic(err)
		}

		totalRisks = len(risks)

		areaPriority, err = models.GetPriorityByProjectId(uint(projectId))
		if err != nil {
			panic(err)
		}
	}

	return c.Render("contoh", fiber.Map{
		"TotalAssets":  totalAssets,
		"TotalRisks":   totalRisks,
		"Projects":     projects,
		"AreaPriority": areaPriority,
		"ProjectId":    selected,
		"Category1":    category1,
		"Category2":    category2,
		"Category3":    category3,
		"Category4":    category4,
		"Risks":        risks,
		"Assets":       assets,
		"Selected":     selected,
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

type FormRequest struct {
	Title             string `form:"title"`
	Area1             string `form:"area1"`
	Area2             string `form:"area2"`
	Area3             string `form:"area3"`
	Area4             string `form:"area4"`
	Area5             string `form:"area5"`
	CriticalAsset     string `form:"critical_asset"`
	Rationale         string `form:"rationale"`
	AssetDesc         string `form:"asset_desc"`
	AssetOwners       string `form:"asset_owners"`
	Confidentiality   string `form:"confidentiality"`
	Integrity         string `form:"integrity"`
	Availability      string `form:"availability"`
	Misr              string `form:"misr"`
	ContainerOwners   string `form:"container_owners"`
	TechnicalInternal string `form:"technical_internal"`
	TechnicalExternal string `form:"technical_external"`
	PhysicalInternal  string `form:"physical_internal"`
	PhysicalExternal  string `form:"physical_external"`
	PeopleInternal    string `form:"people_internal"`
	PeopleExternal    string `form:"people_external"`
}

func FormPost(c *fiber.Ctx) error {
	r := new(FormRequest)
	err := c.BodyParser(r)
	if err != nil {
		panic(err)
	}

	// Insert Project
	projectId, err := models.InsertProject(r.Title)
	if err != nil {
		panic(err)
	}

	area1, err := strconv.Atoi(r.Area1)
	if err != nil {
		panic(err)
	}

	area2, err := strconv.Atoi(r.Area2)
	if err != nil {
		return err
	}

	area3, err := strconv.Atoi(r.Area3)
	if err != nil {
		return err
	}

	area4, err := strconv.Atoi(r.Area4)
	if err != nil {
		return err
	}

	area5, err := strconv.Atoi(r.Area5)
	if err != nil {
		return err
	}

	// Insert Impact Area Priority
	_, err = models.InsertPriority(projectId, area1, area2, area3, area4, area5)
	if err != nil {
		return err
	}

	// Insert Critical Asset Information
	assetId, err := models.InsertAssetProfile(projectId, r.CriticalAsset, r.Rationale, r.AssetDesc, r.AssetOwners, r.Confidentiality, r.Integrity, r.Availability, r.Misr)
	if err != nil {
		return err
	}

	// Insert Container
	_, err = models.InsertContainer(projectId, assetId, r.ContainerOwners, r.TechnicalInternal, r.TechnicalExternal, r.PhysicalInternal, r.PhysicalExternal, r.PeopleInternal, r.PeopleExternal)
	if err != nil {
		return err
	}

	return c.SendString(`<div class="overlay" id="overlay"><div class="notification-box"><h4>Success</h4><a href="/notif">ok!</a></div></div>`)
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
