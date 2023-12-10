package models

import (
	"final_allegro/database"
)

type Project struct {
	ID    uint   `gorm:"primaryKey"`
	Title string `gorm:"not null"`
}

func InsertProject(title string) (uint, error) {
	project := Project{Title: title}
	result := database.DB.Create(&project)

	if result.Error != nil {
		return 0, result.Error
	}

	return project.ID, nil
}

func GetProjectById(projectId uint) (Project, error) {
	var project Project
	result := database.DB.First(&project, projectId)
	return project, result.Error
}

func GetAllProjects() ([]Project, error) {
	var projects []Project
	result := database.DB.Find(&projects)
	return projects, result.Error
}

func DeleteProjectById(projectId uint) error {
	prior, err := GetPriorityByProjectId(projectId)
	if err != nil {
		return err
	}

	// Delete Prior
	result := database.DB.Delete(&ImpactPriority{}, prior.ID)
	if result.Error != nil {
		return result.Error
	}

	assets, err := GetAllAssetByProjectId(projectId)
	if err != nil {
		return err
	}

	risks, err := GetAllRisksByProjectId(projectId)
	if err != nil {
		return err
	}

	mitigations, err := GetAllMitigationByProjectId(projectId)
	if err != nil {
		return err
	}

	containers, err := GetAllContainersByProjectId(projectId)
	if err != nil {
		return err
	}

	// Delete Container
	for _, container := range containers {
		result = database.DB.Delete(&AssetContainer{}, container.ID)
		if result.Error != nil {
			return result.Error
		}
	}

	// Delete Mitgation
	for _, mitigation := range mitigations {
		result = database.DB.Delete(&RiskMitigation{}, mitigation.ID)
		if result.Error != nil {
			return result.Error
		}
	}

	// Delete Risk
	for _, risk := range risks {
		result = database.DB.Delete(&RiskMitigation{}, risk.ID)
		if result.Error != nil {
			return result.Error
		}
	}

	// Delete Asset
	for _, asset := range assets {
		result := database.DB.Delete(&AssetInformation{}, asset.ID)
		if result.Error != nil {
			return result.Error
		}
	}

	// Delete Project
	result = database.DB.Delete(&Project{}, projectId)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
