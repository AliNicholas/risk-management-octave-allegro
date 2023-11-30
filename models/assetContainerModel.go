package models

type AssetContainer struct {
	ID        uint
	AssetID   uint   `gorm:"not null"`
	ProjectID uint   `gorm:"not null"`
	Owners    string `gorm:"not null"`
	Technical string
	Physical  string
	People    string
}
