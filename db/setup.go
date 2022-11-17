package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func init() {
	var err error
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	// Logger config for queries set to level 4 (Info)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info, // Log level Info will output everything
		},
	)

	// Establish connection to the db
	DB, err = gorm.Open(postgres.Open(os.Getenv("DB_SOURCE")), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Fatal("unable to connect to the database:", err)
	}

	fmt.Println("Successfully conncected to the database")

	err = DB.AutoMigrate(&User{}, &Project{}, &Team{}, &Role{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Migration successfull")

	// TODO: More config for the DB
}

func GetDB() *gorm.DB {
	return DB
}
