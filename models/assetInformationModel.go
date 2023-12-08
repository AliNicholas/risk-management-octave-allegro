package models

import "final_allegro/database"

type AssetInformation struct {
	ID                               uint   `gorm:"primaryKey"`
	ProjectID                        uint   `gorm:"not null"`
	CriticalAsset                    string `gorm:"not null"`
	RationaleForSelection            string `gorm:"not null"`
	Description                      string `gorm:"not null"`
	Owners                           string `gorm:"not null"`
	Confidentiality                  string `gorm:"not null"`
	Integrity                        string `gorm:"not null"`
	Availability                     string `gorm:"not null"`
	MostImportantSecurityRequirement string `gorm:"not null"`

	Project Project `gorm:"foreignKey:ProjectID"`
}

func InsertAssetProfile(projectId uint, criticalAsset, rationaleForSelection, description, owners, confidentiality, integrity, availability, mostImportantSecurityRequirement string) (uint, error) {
	assetInformation := AssetInformation{
		ProjectID:                        projectId,
		CriticalAsset:                    criticalAsset,
		RationaleForSelection:            rationaleForSelection,
		Description:                      description,
		Owners:                           owners,
		Confidentiality:                  confidentiality,
		Integrity:                        integrity,
		Availability:                     availability,
		MostImportantSecurityRequirement: mostImportantSecurityRequirement,
	}

	result := database.DB.Create(&assetInformation)

	if result.Error != nil {
		return 0, result.Error
	}

	return assetInformation.ID, nil
}

func GetAssetInformationById(id uint) (AssetInformation, error) {
	var assetInformation AssetInformation
	result := database.DB.First(&assetInformation, id)
	return assetInformation, result.Error
}

func GetAllAssetByProjectId(projectID uint) ([]AssetInformation, error) {
	var assetProfiles []AssetInformation
	result := database.DB.Where("project_id = ?", projectID).Find(&assetProfiles)
	return assetProfiles, result.Error
}
