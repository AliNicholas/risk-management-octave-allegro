package models

import "final_allegro/database"

type AssetContainer struct {
	ID                uint   `gorm:"primaryKey"`
	ProjectID         uint   `gorm:"not null"`
	AssetID           uint   `gorm:"not null"`
	Owners            string `gorm:"not null"`
	TechnicalInternal string `gorm:"not null"`
	TechnicalExternal string `gorm:"not null"`
	PhysicalInternal  string `gorm:"not null"`
	PhysicalExternal  string `gorm:"not null"`
	PeopleInternal    string `gorm:"not null"`
	PeopleExternal    string `gorm:"not null"`
}

func InsertContainer(projectID, assetID uint, owners, technicalInternal, technicalExternal, physicalInternal, physicalExternal, peopleInternal, peopleExternal string) (uint, error) {
	container := AssetContainer{
		ProjectID:         projectID,
		AssetID:           assetID,
		Owners:            owners,
		TechnicalInternal: technicalInternal,
		TechnicalExternal: technicalExternal,
		PhysicalInternal:  physicalInternal,
		PhysicalExternal:  physicalExternal,
		PeopleInternal:    peopleInternal,
		PeopleExternal:    peopleExternal,
	}

	result := database.DB.Create(&container)
	if result.Error != nil {
		return 0, result.Error
	}

	return container.ID, nil
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
