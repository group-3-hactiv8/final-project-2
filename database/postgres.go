package database

import (
	"final-project-2/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// const (
// 	host     = "localhost"
// 	user     = "postgres"
// 	password = "postgres"
// 	dbPort   = 5432
// 	dbname   = "final-project-2"
// )

const (
	host     = "containers-us-west-5.railway.app"
	user     = "postgres"
	password = "wOXcQq26wtArgQ6aNqRS"
	dbPort   = 7636
	dbname   = "railway"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		host, user, password, dbname, dbPort,
	)
	dsn := config
	// inject variable db
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	err = db.Debug().AutoMigrate(models.Comment{}, models.Photo{}, models.SocialMedia{}, models.User{})

	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	log.Println("Successfully connected to database")

}

func GetPostgresInstance() *gorm.DB {
	return db
}
