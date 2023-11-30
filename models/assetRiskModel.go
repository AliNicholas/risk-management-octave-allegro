package models

type AssetRisk struct {
	ID                   uint   `gorm:"primaryKey"`
	ContainerID          uint   `gorm:"not null"`
	ProjectID            uint   `gorm:"not null"`
	AreaOfConcern        string `gorm:"not null"`
	Actor                string `gorm:"not null"`
	Means                string `gorm:"not null"`
	Motive               string `gorm:"not null"`
	Outcome              string `gorm:"not null"`
	SecurityRequirements string `gorm:"not null"`
	Probability          string `gorm:"not null"`
	Consequences         string `gorm:"not null"`
	Severity             string `gorm:"not null"`
}

func GetAllProjectAssets() {

}

func GetAllProjectAsse() {

}
