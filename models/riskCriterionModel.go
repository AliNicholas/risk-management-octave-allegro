package models

type RiskCriterion struct {
	ID         uint
	ProjectID  uint   `gorm:"not null"`
	Field      string `gorm:"not null"`
	ImpactArea string `gorm:"not null"`
	Low        string `gorm:"not null"`
	Moderate   string `gorm:"not null"`
	High       string `gorm:"not null"`
}
