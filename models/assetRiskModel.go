package models

import "final_allegro/database"

type AssetRisk struct {
	ID                   uint   `gorm:"primaryKey"`
	ContainerID          uint   `gorm:"not null"`
	ProjectID            uint   `gorm:"not null"`
	AreaOfConcern        string `gorm:"not null"`
	Actor                string `gorm:"not null"`
	Means                string `gorm:"not null"`
	Motive               string `gorm:"not null"`
	Outcome              string `gorm:"not null"`
	SecurityRequirements string `gorm:"not null"`
	Probability          string `gorm:"not null"`
	Consequences         string `gorm:"not null"`
	Severity             string `gorm:"not null"`
}

func InsertRisk(containerId, projectId uint, areaOfConcern, actor, means, motive, outcome, securityRequirements, probability, consequences, severity string) error {
	risk := AssetRisk{
		ContainerID:          containerId,
		ProjectID:            projectId,
		AreaOfConcern:        areaOfConcern,
		Actor:                actor,
		Means:                means,
		Motive:               motive,
		Outcome:              outcome,
		SecurityRequirements: securityRequirements,
		Probability:          probability,
		Consequences:         consequences,
		Severity:             severity,
	}

	result := database.DB.Create(&risk)
	return result.Error
}

func InsertRiskBatch(risks []AssetRisk) error {
	result := database.DB.Create(&risks)
	return result.Error
}

func GetRiskById(id uint) (AssetRisk, error) {
	var risk AssetRisk
	result := database.DB.First(&risk, id)
	return risk, result.Error
}

func GetAllRisksByContainerId(containerId uint) ([]AssetRisk, error) {
	var risks []AssetRisk
	result := database.DB.Where("container_id = ?", containerId).Find(&risks)
	return risks, result.Error
}

func GetAllRisksByProjectId(projectId uint) ([]AssetRisk, error) {
	var risks []AssetRisk
	result := database.DB.Where("project_id = ?", projectId).Find(&risks)
	return risks, result.Error
}
