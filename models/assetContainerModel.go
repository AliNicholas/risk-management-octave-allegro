package models

import "final_allegro/database"

type AssetContainer struct {
	ID        uint   `gorm:"primaryKey"`
	AssetID   uint   `gorm:"not null"`
	ProjectID uint   `gorm:"not null"`
	Owners    string `gorm:"not null"`
	Technical string `gorm:"not null"`
	Physical  string `gorm:"not null"`
	People    string `gorm:"not null"`
}

func InsertContainer(assetID, projectID uint, owners, technical, physical, people string) error {
	container := AssetContainer{
		AssetID:   assetID,
		ProjectID: projectID,
		Owners:    owners,
		Technical: technical,
		Physical:  physical,
		People:    people,
	}

	result := database.DB.Create(&container)
	return result.Error
}

func InsertContainerBatch(containers []AssetContainer) error {
	result := database.DB.Create(&containers)
	return result.Error
}

func GetContainerById(id uint) (AssetContainer, error) {
	var container AssetContainer
	result := database.DB.First(&container, id)
	return container, result.Error
}

func GetAllContainersByAssetId(assetId uint) ([]AssetContainer, error) {
	var containers []AssetContainer
	result := database.DB.Where("asset_id = ?", assetId).Find(&containers)
	return containers, result.Error
}

func GetAllContainersByProjectId(projectId uint) ([]AssetContainer, error) {
	var containers []AssetContainer
	result := database.DB.Where("project_id = ?", projectId).Find(&containers)
	return containers, result.Error
}
