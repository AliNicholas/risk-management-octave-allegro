package main

import (
	"final_allegro/database"
	"final_allegro/models"
)

func main() {
	database.ConnectDB()

	database.DB.Debug().AutoMigrate(
		&models.Project{},
		&models.ImpactPriority{},
		&models.AssetInformation{},
		&models.AssetContainer{},
		&models.AssetRisk{},
		&models.RiskMitigation{},
	)

}
