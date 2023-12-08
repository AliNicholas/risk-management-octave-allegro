package models

import "final_allegro/database"

type AssetRisk struct {
	ID                   uint   `gorm:"primaryKey"`
	ProjectID            uint   `gorm:"not null"`
	AssetID              uint   `gorm:"not null"`
	AreaOfConcern        string `gorm:"not null"`
	Actor                string `gorm:"not null"`
	Means                string `gorm:"not null"`
	Motive               string `gorm:"not null"`
	Outcome              string `gorm:"not null"`
	SecurityRequirements string `gorm:"not null"`
	Probability          string `gorm:"not null"`
	Consequences         string `gorm:"not null"`
	ReputationConfidence string `gorm:"not null"`
	Financial            string `gorm:"not null"`
	Productivity         string `gorm:"not null"`
	SafetyHealth         string `gorm:"not null"`
	FinesLegalPenalties  string `gorm:"not null"`

	Project          Project          `gorm:"foreignKey:ProjectID"`
	AssetInformation AssetInformation `gorm:"foreignKey:AssetID"`
}

func InsertRisk(risk AssetRisk) (uint, error) {

	result := database.DB.Create(&risk)
	if result.Error != nil {
		return 0, result.Error
	}
	return risk.ID, nil
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
