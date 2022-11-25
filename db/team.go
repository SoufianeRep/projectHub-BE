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
func CreateTeam(arg CreateTeamParams) (Team, error) {
	team := Team{
		TeamName: arg.TeamName,
	}
	result := DB.Create(&team)
	if result.Error != nil {
		return Team{}, result.Error
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

// GetProjects takes a team as a receiver and returns all the projects belonging to that team and an error.
func (team Team) GetProjects() ([]Project, error) {
	projects := []Project{}

	rows, err := DB.Table("projects").
		Joins("JOIN teams ON projects.team_id = teams.id").
		Where("projects.team_id = ?", team.ID).
		Rows()
	if err != nil {
		return []Project{}, err
	}

	for rows.Next() {
		DB.ScanRows(rows, &projects)
	}

	return projects, nil
}

// GetMembers takes a team as a receiver and return a list of its members
func (team Team) GetMembers() ([]User, error) {
	users := []User{}

	rows, err := DB.Table("users").
		Joins("JOIN roles ON roles.user_id = users.id").
		Joins("JOIN teams On roles.team_id = teams.id").
		Where("teams.id = ?", team.ID).
		Rows()
	if err != nil {
		return []User{}, err
	}

	for rows.Next() {
		DB.ScanRows(rows, &users)
	}

	return users, nil
}
