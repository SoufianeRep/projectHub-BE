package db

import (
	"fmt"

	"gorm.io/gorm"
)

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

	rows, err := DB.Model(Project{}).
		Where("projects.team_id = ?", team.ID).
		Joins("JOIN teams ON projects.team_id = teams.id").
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

	rows, err := DB.Model(User{}).
		Where("teams.id = ?", team.ID).
		Joins("JOIN roles ON roles.user_id = users.id").
		Joins("JOIN teams On roles.team_id = teams.id").
		Rows()
	if err != nil {
		return []User{}, err
	}

	for rows.Next() {
		DB.ScanRows(rows, &users)
	}

	for _, u := range users {
		fmt.Println(u.ID)
		fmt.Println(u.Email)
	}
	return users, nil
}

// AddMember adds a user to the reciever team with the provided role
func (team Team) AddMember(email string, role string) error {
	user, err := GetUserByEmail(email)
	if err != nil {
		return err
	}

	arg := CreateRoleParams{
		UserID: user.ID,
		TeamID: team.ID,
		Role:   role,
	}

	err = CreateRole(arg)
	if err != nil {
		return err
	}
	return nil
}
