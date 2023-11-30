package models

import "final_allegro/database"

type AssetProfile struct {
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
}

func InsertAssetProfile(projectId uint, criticalAsset, rationaleForSelection, description, owners, confidentiality, integrity, availability, mostImportantSecurityRequirement string) error {
	assetProfile := AssetProfile{
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

	result := database.DB.Create(&assetProfile)
	return result.Error
}

func GetAssetProfileById(id uint) (AssetProfile, error) {
	var assetProfile AssetProfile
	result := database.DB.First(&assetProfile, id)
	return assetProfile, result.Error
}

func GetAssetProfileByProjectId(projectID uint) ([]AssetProfile, error) {
	var assetProfiles []AssetProfile
	result := database.DB.Where("project_id = ?", projectID).Find(&assetProfiles)
	return assetProfiles, result.Error
}
