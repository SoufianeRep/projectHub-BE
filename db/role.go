package db

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Role   string `gorm:"default:manager"`
	UserID uint   `gorm:"not null"`
	User   *User
	TeamID uint `gorm:"not null"`
	Team   *Team
}

type CreateRoleParams struct {
	UserID uint
	TeamID uint
	Role   string `json:"role"`
}

func CreateRole(arg *CreateRoleParams) error {
	role := &Role{
		UserID: arg.UserID,
		TeamID: arg.TeamID,
		Role:   arg.Role,
	}
	result := DB.Create(&role)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
