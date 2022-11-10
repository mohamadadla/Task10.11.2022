package main

import (
	"adla.com/task10.11.2022/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	studentController := controller.StrudentController{}

	// Routes
	e.GET("/students", studentController.GetAll)
	e.POST("/students", studentController.Create)
	e.GET("/students/:id", studentController.Get)
	e.PUT("/students/:id", studentController.PUT)
	e.DELETE("/students/:id", studentController.DELETE)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
