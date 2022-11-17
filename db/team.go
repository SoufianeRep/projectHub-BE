package db

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	TeamName string `gorm:"not null"`
	Projects []Project
}

type CreateTeamParams struct {
	TeamName string `json:"name"`
}

// CreatTeam creates a team with the gven argument name and return it or an error
func CreateTeam(arg CreateTeamParams) (*Team, error) {
	team := &Team{
		TeamName: arg.TeamName,
	}
	result := DB.Create(&team)
	if result.Error != nil {
		return nil, result.Error
	}
	return team, nil
}

// GetTeam return the first team with provided primary key
func GetTeam(id uint) (*Team, error) {
	team := &Team{}

	result := DB.First(team, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return team, nil
}
