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

	projects, err := models.GetAllProjects()
	if err != nil {
		return err
	}

	return c.Render("dashboard", fiber.Map{
		"Projects": projects,
	})
}

func Render(c *fiber.Ctx) error {
	projectIdStr := c.FormValue("project-selector")

	projects, err := models.GetAllProjects()
	if err != nil {
		return err
	}

	var (
		totalAssets  int
		totalRisks   int
		category1    int
		category2    int
		category3    int
		category4    int
		projectId    int
		assets       []models.AssetInformation
		risks        []models.AssetRisk
		areaPriority models.ImpactPriority
	)

	var isTable bool = false

	if projectIdStr != "NONE" {
		projectId, err = strconv.Atoi(projectIdStr)
		if err != nil {
			return err
		}

		assets, err = models.GetAllAssetByProjectId(uint(projectId))
		if err != nil {
			return err
		}

		totalAssets = len(assets)

		risks, err = models.GetAllRisksByProjectId(uint(projectId))
		if err != nil {
			return err
		}

		totalRisks = len(risks)

		areaPriority, err = models.GetPriorityByProjectId(uint(projectId))
		if err != nil {
			return err
		}
		isTable = true
		// isContent = true
	}

	data := map[string]interface{}{
		"TotalAssets": totalAssets,
		"TotalRisks":  totalRisks,
		"Projects":    projects,
		"Category1":   category1,
		"Category2":   category2,
		"Category3":   category3,
		"Category4":   category4,
		"Risks":       risks,
		"Assets":      assets,
	}

	if isTable {
		data["AreaPriority"] = areaPriority
	}

	if projectId > 0 {
		data["ProjectId"] = projectId
	}

	return c.Render("renderDashboard", data)
}

func RenderRiskContent(c *fiber.Ctx) error {
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

	if projectIdStr != "NONE" {
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

	return c.Render("renderDashboard", fiber.Map{
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
func DeleteProject(c *fiber.Ctx) error {
	projectIdStr := c.Params("projectId")

	projectId, err := strconv.Atoi(projectIdStr)
	if err != nil {
		return err
	}

	err = models.DeleteProjectById(uint(projectId))
	if err != nil {
		return err
	}

	return c.SendString(`<div class="overlay" id="overlay"><div class="notification-box"><h4>Success</h4><a href="/notif">ok!</a></div></div>`)
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
func AddProfilePage(c *fiber.Ctx) error {
	projectId := c.Params("projectId")

	return c.Render("add-profile", fiber.Map{
		"ProjectId": projectId,
	})
}

type AddProfileRequest struct {
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

func PostProfile(c *fiber.Ctx) error {
	projectIdStr := c.Params("projectId")
	fmt.Println(projectIdStr)

	projectId, err := strconv.Atoi(projectIdStr)
	if err != nil {
		return err
	}

	r := new(AddProfileRequest)
	err = c.BodyParser(r)
	if err != nil {
		return err
	}

	assetId, err := models.InsertAssetProfile(uint(projectId), r.CriticalAsset, r.Rationale, r.AssetDesc, r.AssetOwners, r.Confidentiality, r.Integrity, r.Availability, r.Misr)
	if err != nil {
		return err
	}

	_, err = models.InsertContainer(uint(projectId), assetId, r.ContainerOwners, r.TechnicalInternal, r.TechnicalExternal, r.PhysicalInternal, r.PhysicalExternal, r.PeopleInternal, r.PeopleExternal)
	if err != nil {
		return err
	}

	return c.SendString(`<div class="overlay" id="overlay"><div class="notification-box"><h4>Success</h4><a href="/notif">ok!</a></div></div>`)
}

// Risk Page
func AddRiskPage(c *fiber.Ctx) error {
	projectIdStr := c.Params("projectId")

	projectId, err := strconv.Atoi(projectIdStr)
	if err != nil {
		return err
	}

	assets, err := models.GetAllAssetByProjectId(uint(projectId))
	if err != nil {
		return err
	}

	return c.Render("add-risk", fiber.Map{
		"Assets":    assets,
		"ProjectId": projectId,
	})
}

func RenderAsset(c *fiber.Ctx) error {
	assetIdStr := c.FormValue("asset-selector")

	assetId, err := strconv.Atoi(assetIdStr)
	if err != nil {
		return err
	}

	asset, err := models.GetAssetInformationById(uint(assetId))
	if err != nil {
		return err
	}

	container, err := models.GetContainerByAssetId(uint(assetId))
	if err != nil {
		return err
	}

	return c.Render("renderAsset", fiber.Map{
		"Asset":     asset,
		"Container": container,
	})
}

type AddRiskRequest struct {
	AssetId             string `form:"asset_id"`
	Concern             string `form:"concern"`
	Actor               string `form:"actor"`
	Means               string `form:"means"`
	Motive              string `form:"motive"`
	Outcome             string `form:"outcome"`
	SecurityReq         string `form:"security_req"`
	Probability         string `form:"probability"`
	Consequences        string `form:"consequences"`
	Area1               string `form:"area1"`
	Area2               string `form:"area2"`
	Area3               string `form:"area3"`
	Area4               string `form:"area4"`
	Area5               string `form:"area5"`
	TechnicalMitigation string `form:"technical_mitigation"`
	PhysicalMitigation  string `form:"physical_mitigation"`
	PeopleMitigation    string `form:"people_mitigation"`
}

func PostRisk(c *fiber.Ctx) error {
	projectIdStr := c.Params("projectId")

	projectId, err := strconv.Atoi(projectIdStr)
	if err != nil {
		return err
	}

	r := new(AddRiskRequest)
	err = c.BodyParser(r)
	if err != nil {
		return err
	}

	assetId, err := strconv.Atoi(r.AssetId)
	if err != nil {
		fmt.Println(assetId)
		fmt.Println("atas")
		fmt.Println(r.AssetId)
		fmt.Println("bawah")
		return err
	}

	riskId, err := models.InsertRisk(models.AssetRisk{ProjectID: uint(projectId), AssetID: uint(assetId), AreaOfConcern: r.Concern, Actor: r.Actor, Means: r.Means, Motive: r.Motive, Outcome: r.Outcome, SecurityRequirements: r.SecurityReq, Probability: r.Probability, Consequences: r.Consequences, ReputationConfidence: r.Area1, Financial: r.Area2, Productivity: r.Area3, SafetyHealth: r.Area4, FinesLegalPenalties: r.Area5})
	if err != nil {
		return err
	}

	_, err = models.InsertRiskMitigation(models.RiskMitigation{ProjectID: uint(projectId), RiskID: riskId, TechnicalMitigation: r.TechnicalMitigation, PhysicalMitigation: r.PhysicalMitigation, PeopleMitigation: r.PeopleMitigation})
	if err != nil {
		return err
	}

	return c.SendString(`<div class="overlay" id="overlay"><div class="notification-box"><h4>Success</h4><a href="/notif">ok!</a></div></div>`)
}
func HandleScore(c *fiber.Ctx) error {
	projectIdStr := c.Params("projectId")
	areaStr := c.Params("area")
	recentScoreStr := c.Params("recentScore")
	totalScoreStr := c.FormValue("total_score")

	var newScore int
	var multiplier int

	projectId, err := strconv.Atoi(projectIdStr)
	if err != nil {
		return err
	}
	area, err := strconv.Atoi(areaStr)
	if err != nil {
		return err
	}
	recentscore, err := strconv.Atoi(recentScoreStr)
	if err != nil {
		return err
	}
	totalScore, err := strconv.Atoi(totalScoreStr)
	if err != nil {
		return err
	}
	priority, err := models.GetPriorityByProjectId(uint(projectId))
	if err != nil {
		return err
	}

	if area == 1 {
		multiplierStr := c.FormValue("area1")
		multiplier, err = strconv.Atoi(multiplierStr)
		if err != nil {
			return err
		}

		totalScore -= recentscore
		newScore = multiplier * priority.ReputationConfidence
		totalScore += newScore

	} else if area == 2 {
		multiplierStr := c.FormValue("area2")
		multiplier, err = strconv.Atoi(multiplierStr)
		if err != nil {
			return err
		}
		totalScore -= recentscore
		newScore = multiplier * priority.Financial
		totalScore += newScore

	} else if area == 3 {
		multiplierStr := c.FormValue("area3")
		multiplier, err = strconv.Atoi(multiplierStr)
		if err != nil {
			return err
		}

		totalScore -= recentscore
		newScore = multiplier * priority.Productivity
		totalScore += newScore

	} else if area == 4 {
		multiplierStr := c.FormValue("area4")
		multiplier, err = strconv.Atoi(multiplierStr)
		if err != nil {
			return err
		}

		totalScore -= recentscore
		newScore = multiplier * priority.SafetyHealth
		totalScore += newScore

	} else if area == 5 {
		multiplierStr := c.FormValue("area5")
		multiplier, err = strconv.Atoi(multiplierStr)
		if err != nil {
			return err
		}

		totalScore -= recentscore
		newScore = multiplier * priority.FinesLegalPenalties
		totalScore += newScore
	}

	return c.Render("handleScore", fiber.Map{
		"TotalScore": totalScore,
		"NewScore":   newScore,
		"Area":       area,
		"Selected":   multiplier,
		"ProjectId":  projectId,
	})
}
