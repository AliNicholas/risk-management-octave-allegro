package models

import "final_allegro/database"

type ImpactPriority struct {
	ID                   uint `gorm:"primaryKey"`
	ProjectID            uint `gorm:"not null"`
	ReputationConfidence int  `gorm:"not null"`
	Financial            int  `gorm:"not null"`
	Productivity         int  `gorm:"not null"`
	SafetyHealth         int  `gorm:"not null"`
	FinesLegalPenalties  int  `gorm:"not null"`
}

func InsertPriority(projectId uint, reputationConfidence, financial, productivity, safetyHealth, finesLegalPenalties int) error {
	priority := ImpactPriority{
		ProjectID:            projectId,
		ReputationConfidence: reputationConfidence,
		Financial:            financial,
		Productivity:         productivity,
		SafetyHealth:         safetyHealth,
		FinesLegalPenalties:  finesLegalPenalties,
	}

	result := database.DB.Create(&priority)
	return result.Error
}

func GetPriorityById(id uint) (ImpactPriority, error) {
	var priority ImpactPriority
	result := database.DB.First(&priority, id)
	return priority, result.Error
}

func GetPriorityByProjectId(projectId uint) (ImpactPriority, error) {
	var priority ImpactPriority
	result := database.DB.Where("project_id = ?", projectId).First(&priority)
	return priority, result.Error
}
