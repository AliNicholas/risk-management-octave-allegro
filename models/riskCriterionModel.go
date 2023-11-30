package models

import "final_allegro/database"

type RiskCriterion struct {
	ID         uint   `gorm:"primaryKey"`
	ProjectID  uint   `gorm:"not null"`
	Field      string `gorm:"not null"`
	ImpactArea string `gorm:"not null"`
	Low        string `gorm:"not null"`
	Moderate   string `gorm:"not null"`
	High       string `gorm:"not null"`
}

func InsertCriterion(projectId uint, field, impactArea, low, moderate, high string) error {
	criterion := RiskCriterion{
		ProjectID:  projectId,
		Field:      field,
		ImpactArea: impactArea,
		Low:        low,
		Moderate:   moderate,
		High:       high,
	}

	result := database.DB.Create(&criterion)
	return result.Error
}

func GetCriterionByProjectId(projectIid uint) (RiskCriterion, error) {
	var criteria RiskCriterion
	result := database.DB.Where("project_id = ?", projectIid).Find(&criteria)
	return criteria, result.Error
}
