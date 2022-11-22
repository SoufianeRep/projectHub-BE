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

func GetUserByEmail(email string) (User, error) {
	user := User{}
	result := DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return User{}, result.Error
	}

	return user, nil
}

func GetUSerByID(id uint) (User, error) {
	user := User{}

	result := DB.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return User{}, result.Error
	}

	return user, nil
}

type UpdateUserParams struct {
	Email      string
	Password   string
	LastSignin time.Time
}

// UpdateLastSign takes a user row instance as a receiver and updates the
// last_signin column to time.Now()
func (user User) UpdateLastSignin() {
	DB.Model(&user).Update("last_signin", time.Now())
}
