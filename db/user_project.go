package db

import "gorm.io/gorm"

type UserProject struct {
	gorm.Model
	UserID    uint    `gorm:""`
	User      User    `gorm:""`
	ProjectID uint    `gorm:""`
	Project   Project `gorm:""`
}
