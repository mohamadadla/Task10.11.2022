package container

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func GetDbConnection() *gorm.DB {

	// Connect to the "student" database
	dsn := "host=localhost user=postgres password=@dla port=3307 database=student_db"
	dbConnection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error configuring the database: ", err)
	}
	log.Println("Hey! You successfully connected to your CockroachDB cluster.")

	return dbConnection
}
