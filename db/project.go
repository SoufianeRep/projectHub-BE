package db

import (
	"github.com/jackc/pgtype"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name      string       `gorm:"not null" json:"name"`
	Language  string       `gorm:"not null" json:"language"`
	Length    uint         `gorm:"not null" json:"length"`
	Trancript pgtype.JSONB `gorm:"type:jsonb" json:"transcript"`
	TeamID    uint         `gorm:"not null"`
	Team      Team
}

type CreatePRojectParams struct {
	Name     string `json:"name"`
	Language string `json:"language"`
	Length   string `json:"length"`
}
