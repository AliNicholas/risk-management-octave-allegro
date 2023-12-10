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

func DeleteAssetByAssetId(assetId uint) error {
	container, err := GetContainerByAssetId(assetId)
	if err != nil {
		return err
	}

	risks, err := GetAllRisksByAssetId(assetId)
	if err != nil {
		return err
	}

	// Delete Container
	result := database.DB.Delete(&AssetContainer{}, container.ID)
	if result.Error != nil {
		return result.Error
	}

	// Delete Risk
	for _, risk := range risks {

		// Delete Mitigation
		mitigation, err := GetMitigationByRiskId(risk.ID)
		if err != nil {
			return err
		}

		result = database.DB.Delete(&RiskMitigation{}, mitigation.ID)
		if result.Error != nil {
			return result.Error
		}

		result := database.DB.Delete(&AssetRisk{}, risk.ID)
		if result.Error != nil {
			return result.Error
		}
	}

	// Delete Asset
	result = database.DB.Delete(&AssetInformation{}, assetId)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
