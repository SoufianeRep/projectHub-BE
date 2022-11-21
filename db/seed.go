package db

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/SoufianeRep/tscit/util"
	"gorm.io/gorm"
)

// RunSeed seeds the provided db with basic seeds for dev purposes
func RunSeed(db *gorm.DB) error {
	users, err := seedUsers(db)
	if err != nil {
		return err
	}
	teams, err := seedTeams(db)
	if err != nil {
		return err
	}
	err = seedRoles(db, users, teams)
	if err != nil {
		return err
	}

	_, err = seedProjects(db, teams)
	if err != nil {
		return err
	}
	return nil
}

func seedUsers(db *gorm.DB) (users []User, err error) {
	n := 10

	for i := 0; i < n; i++ {
		user := User{
			Email:      util.RandomEmail(),
			Password:   "password",
			LastSignin: time.Now(),
		}

		result := db.Create(&user)
		if result.Error != nil {
			err = fmt.Errorf("problem has occured while seeding users:", result.Error)
		}

		users = append(users, user)
	}
	return
}

func seedTeams(db *gorm.DB) (teams []Team, err error) {
	n := 2

	for i := 0; i < n; i++ {
		team := Team{
			TeamName: util.RandomString(8),
		}
		result := db.Create(&team)
		if result.Error != nil {
			err = fmt.Errorf("roblem has occured while seeding Team: %v", result.Error)
		}
		teams = append(teams, team)
	}
	return
}

func seedRoles(db *gorm.DB, users []User, teams []Team) (err error) {
	n := len(users)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		role := Role{
			Role:   util.RandomRole(),
			UserID: users[i].ID,
			TeamID: teams[rand.Intn(2)].ID,
		}

		result := db.Create(&role)
		if result.Error != nil {
			err = fmt.Errorf("problem has occured while seeding roles: %v", result.Error)
		}
	}
	return
}

func seedProjects(db *gorm.DB, teams []Team) (projects []Project, err error) {
	n := 20
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		project := Project{
			Name:     util.RandomString(10),
			Language: "english",
			Length:   uint(util.RandomInt(5, 25)),
			TeamID:   teams[rand.Intn(2)].ID,
		}
		err = project.Transcript.Set([]string{})
		if err != nil {
			return nil, err
		}

		result := db.Create(&project)
		if result.Error != nil {
			err = fmt.Errorf("a problem has occured while seeding projects: %v", result.Error)
		}
	}
	return
}
