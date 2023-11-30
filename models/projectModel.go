package models

import "final_allegro/database"

type Project struct {
	ID    uint   `gorm:"primaryKey"`
	Title string `gorm:"not null"`
}

func InsertProject(title string) error {
	project := Project{Title: title}
	result := database.DB.Create(&project)
	return result.Error
}

func GetAllProjects() ([]Project, error) {
	var projects []Project
	result := database.DB.Find(&projects)
	return projects, result.Error
}
