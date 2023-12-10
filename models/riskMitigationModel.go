package models

import "final_allegro/database"

type RiskMitigation struct {
	ID                  uint   `gorm:"primaryKey"`
	ProjectID           uint   `gorm:"not null"`
	RiskID              uint   `gorm:"not null"`
	TechnicalMitigation string `gorm:"not null"`
	PhysicalMitigation  string `gorm:"not null"`
	PeopleMitigation    string `gorm:"not null"`

	Project   Project   `gorm:"foreignKey:ProjectID"`
	AssetRisk AssetRisk `gorm:"foreignKey:RiskID"`
}

func InsertRiskMitigation(riskMitigation RiskMitigation) (uint, error) {
	result := database.DB.Create(&riskMitigation)
	if result.Error != nil {
		return 0, result.Error
	}

	return riskMitigation.ID, nil
}

func GetMitigationById(id uint) (RiskMitigation, error) {
	var riskMitigation RiskMitigation
	result := database.DB.First(&riskMitigation, id)
	return riskMitigation, result.Error
}

func GetAllMitigationByProjectId(projectId uint) ([]RiskMitigation, error) {
	var riskMitigations []RiskMitigation
	result := database.DB.Where("project_id = ?", projectId).Find(&riskMitigations)
	return riskMitigations, result.Error
}

func GetMitigationByRiskId(riskId uint) (RiskMitigation, error) {
	var riskMitigations RiskMitigation
	result := database.DB.Where("risk_id = ?", riskId).Find(&riskMitigations)
	return riskMitigations, result.Error
}
