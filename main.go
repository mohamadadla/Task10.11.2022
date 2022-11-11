package main

import (
	"adla.com/task10.11.2022/controller"
	"adla.com/task10.11.2022/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Connect to the "student" database
	dsn := "host=localhost user=postgres password=@dla port=3307"
	dbConnection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error configuring the database: ", err)
	}
	log.Println("Hey! You successfully connected to your CockroachDB cluster.")

	studentController := controller.StudentController{
		StudentRepository: repository.StudentRepository{
			DB: dbConnection,
		},
	}

	// Routes
	e.GET("/student", studentController.GetAll)
	e.POST("/student", studentController.Create)
	e.GET("/student/:id", studentController.Get)
	e.PUT("/student/:id", studentController.PUT)
	e.DELETE("/student/:id", studentController.DELETE)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
