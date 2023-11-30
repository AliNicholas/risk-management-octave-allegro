package models


type ImpactPriority struct {
	ID uint
	ProjectID              uint `gorm:"not null"`
	ReputationConfidence  int  `gorm:"not null"`
	Financial             int  `gorm:"not null"`
	Productivity          int  `gorm:"not null"`
	SafetyHealth          int  `gorm:"not null"`
	FinesLegalPenalties   int  `gorm:"not null"`
}
