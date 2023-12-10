package models

import (
	"final_allegro/database"
	"final_allegro/helper"
	"strconv"
)

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
	RelativeScore        string `gorm:"not null"`

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

func GetAllRisksByAssetId(assetId uint) ([]AssetRisk, error) {
	var risks []AssetRisk
	result := database.DB.Where("asset_id = ?", assetId).Find(&risks)
	return risks, result.Error
}

func GetAllRisksByProjectId(projectId uint) ([]AssetRisk, error) {
	var risks []AssetRisk
	result := database.DB.Where("project_id = ?", projectId).Find(&risks)
	return risks, result.Error
}

func GetCategoryMatrix(projectId uint) ([]int, error) {
	var categoryCounts []int
	var risks []AssetRisk

	results := database.DB.Where("project_id = ?", projectId).Find(&risks)
	if results.Error != nil {
		return nil, results.Error
	}

	categoryCounts = make([]int, 4)

	for _, risk := range risks {
		score, err := strconv.Atoi(risk.RelativeScore)
		if err != nil {
			return categoryCounts, err
		}

		categoryNum, _, err := helper.ScoringMatrix(risk.Probability, score)
		if err != nil {
			return categoryCounts, err
		}

		switch categoryNum {
		case 1:
			categoryCounts[0]++
		case 2:
			categoryCounts[1]++
		case 3:
			categoryCounts[2]++
		case 4:
			categoryCounts[3]++
		}
	}

	return categoryCounts, nil
}

func GetCategoryMatrixByAssetId(assetId uint) ([]int, error) {
	var categoryCounts []int
	var risks []AssetRisk

	results := database.DB.Where("asset_id = ?", assetId).Find(&risks)
	if results.Error != nil {
		return nil, results.Error
	}

	categoryCounts = make([]int, 4)

	for _, risk := range risks {
		score, err := strconv.Atoi(risk.RelativeScore)
		if err != nil {
			return categoryCounts, err
		}

		categoryNum, _, err := helper.ScoringMatrix(risk.Probability, score)
		if err != nil {
			return categoryCounts, err
		}

		switch categoryNum {
		case 1:
			categoryCounts[0]++
		case 2:
			categoryCounts[1]++
		case 3:
			categoryCounts[2]++
		case 4:
			categoryCounts[3]++
		}
	}

	return categoryCounts, nil
}
