package models

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

func GetAssetProfileByProjectId() {

}
