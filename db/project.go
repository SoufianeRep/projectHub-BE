package db

import (
	"fmt"

	"github.com/jackc/pgtype"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name       string       `gorm:"not null" json:"name"`
	Language   string       `gorm:"not null" json:"language"`
	Length     uint         `gorm:"not null" json:"length"`
	Transcript pgtype.JSONB `gorm:"type:jsonb" json:"transcript"`
	TeamID     uint         `gorm:"not null"`
	Team       Team
}

type CreateProjectParams struct {
	Name     string `json:"name"`
	Language string `json:"language"`
	Length   uint   `json:"length"`
	TeamID   uint   `json:"team_id"`
}

func CreateProject(arg CreateProjectParams) (*Project, error) {
	project := &Project{
		Name:     arg.Name,
		Length:   arg.Length,
		Language: arg.Language,
		TeamID:   arg.TeamID,
	}

	err := project.Transcript.Set([]string{})
	if err != nil {
		return nil, err
	}

	result := DB.Create(&project)
	if result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}

	return project, nil
}

func GetProject(id uint) (Project, error) {
	project := Project{}

	result := DB.Where("id = ?", id).First(&project)
	if result.Error != nil {
		return Project{}, result.Error
	}

	return project, nil
}
