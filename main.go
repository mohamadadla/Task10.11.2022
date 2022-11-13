package main

import (
	"adla.com/task10.11.2022/container"
	"adla.com/task10.11.2022/internal/adapter/db/repository"
	"adla.com/task10.11.2022/internal/interface/api"
	"adla.com/task10.11.2022/internal/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db := container.GetDbConnection()

	studentController := api.StudentController{
		GetAllStudents: usecase.GetAllStudents{
			RepositoryStudent: repository.StudentRepository{
				DB: db,
			},
		},
		FindById: usecase.FindStudentById{
			RepositoryStudent: repository.StudentRepository{
				DB: db,
			},
		},
		CreateStudent: usecase.CreateStudent{
			RepositoryStudent: repository.StudentRepository{
				DB: db,
			},
		},

		DeleteStudent: usecase.DeleteStudent{
			RepositoryStudent: repository.StudentRepository{
				DB: db,
			},
		},

		EditStudent: usecase.EditStudent{
			RepositoryStudent: repository.StudentRepository{
				DB: db,
			},
		},
	}

	// Routes
	e.GET("/student", studentController.GetAll)
	e.GET("/studentId", studentController.Get)
	e.POST("/student", studentController.Create)
	e.PUT("/student", studentController.Put)
	e.DELETE("/studentdel", studentController.Delete)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
