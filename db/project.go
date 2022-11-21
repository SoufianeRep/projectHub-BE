package db

import (
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

	result := DB.Create(&project)
	if result.Error != nil {
		return nil, result.Error
	}

	return project, nil
}
