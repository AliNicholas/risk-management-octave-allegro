package models

import "final_allegro/database"

type Project struct {
	ID    uint   `gorm:"primaryKey"`
	Title string `gorm:"not null"`
}

func InsertProject(title string) (uint, error) {
	project := Project{Title: title}
	result := database.DB.Create(&project)

	if result.Error != nil {
		return 0, result.Error
	}

	return project.ID, nil
}

func GetProjectById(projectId uint) (Project, error) {
	var project Project
	result := database.DB.First(&project, projectId)
	return project, result.Error
}

func GetAllProjects() ([]Project, error) {
	var projects []Project
	result := database.DB.Find(&projects)
	return projects, result.Error
}

func DeleteProjectById(id uint) error {
	result := database.DB.Delete(&Project{}, id)
	return result.Error
}
