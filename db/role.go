package db

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Role   string `gorm:"default:manager"`
	UserID uint   `gorm:"index:idx_role,not null"`
	User   *User
	TeamID uint `gorm:"index:idx_role,not null"`
	Team   *Team
}

type CreateRoleParams struct {
	UserID uint
	TeamID uint
	Role   string `json:"role" binding:"oneof=manager superuser user"`
}

func CreateRole(arg CreateRoleParams) error {
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
