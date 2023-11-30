package models

type AssetContainer struct {
	ID        uint   `gorm:"primaryKey"`
	AssetID   uint   `gorm:"not null"`
	ProjectID uint   `gorm:"not null"`
	Owners    string `gorm:"not null"`
	Technical string `gorm:"not null"`
	Physical  string `gorm:"not null"`
	People    string `gorm:"not null"`
}

func GetAllProjectContainers() {

}

func GetAllAssetContainers() {

}
