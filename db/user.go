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

func CreateUser(arg CreateUserParams) error {
	result := DB.Create(&User{
		Email:      arg.Email,
		Password:   arg.Password,
		LastSignin: time.Now(),
	})
	if result.Error != nil {
		pgError := result.Error.(*pgconn.PgError)
		if errors.Is(result.Error, pgError) {
			switch pgError.Code {
			case "23505":
				return fmt.Errorf("duplicate, email already exist")
			}
		}
		return fmt.Errorf("an error has occured whie creating the user")
	}

	return nil
}

func GetUserByEmail(email string) (User, error) {
	user := User{}
	result := DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return User{}, result.Error
	}

	return user, nil
}
