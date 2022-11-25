package db

import (
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email      string `gorm:"unique;not null" json:"email"`
	Password   string `gorm:"not null" json:"password"`
	Role       Role
	LastSignin time.Time
}

type CreateUserParams struct {
	Email      string
	Password   string
	LastSignin time.Time
}

func CreateUser(arg CreateUserParams) (User, error) {
	user := User{
		Email:      arg.Email,
		Password:   arg.Password,
		LastSignin: time.Now(),
	}

	result := DB.Create(&user)
	if result.Error != nil {
		pgError := result.Error.(*pgconn.PgError)
		if errors.Is(result.Error, pgError) {
			switch pgError.Code {
			case "23505":
				return User{}, fmt.Errorf("duplicate, email already exist")
			}
		}
		return User{}, fmt.Errorf("an error has occured whie creating the user")
	}

	return user, nil
}

// GetUserByEmail takes an email as a string and retruns the first instance from
// the database and a error
func GetUserByEmail(email string) (User, error) {
	user := User{}
	result := DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return User{}, result.Error
	}

	return user, nil
}

func GetUserByID(id uint) (User, error) {
	user := User{}

	result := DB.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return User{}, result.Error
	}

	return user, nil
}

// UpdateLastSign takes a user row instance as a receiver and updates the
// last_signin column to time.Now()
func (user User) UpdateLastSignin() {
	DB.Model(&user).Update("last_signin", time.Now())
}

// GetTeams takes a user as a receiver queries the DB and return a list of all the teams the user is part of
func (user User) GetTeams() ([]Team, error) {
	teams := []Team{}

	rows, err := DB.Table("teams").
		Where("roles.user_id", user.ID).
		Joins("JOIN roles ON roles.team_id = teams.id").
		Joins("JOIN users ON roles.user_id = users.id").
		Rows()
	if err != nil {
		return teams, fmt.Errorf("teams query error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		DB.ScanRows(rows, &teams)
	}

	return teams, nil
}
