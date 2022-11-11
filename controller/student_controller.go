package controller

import (
	"adla.com/task10.11.2022/repository"
	"net/http"

	"adla.com/task10.11.2022/entity"
	"github.com/labstack/echo/v4"
)

type StudentController struct {
	StudentRepository repository.StudentRepository
}

func (r StudentController) GetAll(c echo.Context) error {

	return c.JSON(http.StatusOK, r.StudentRepository.GetAll())
}

func (r StudentController) Get(c echo.Context) error {

	studententity := &entity.Student{}
	return c.JSON(http.StatusOK, studententity)
}

func (r StudentController) Create(c echo.Context) error {

	studententity := &entity.Student{}

	if err := c.Bind(studententity); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, studententity)
}

func (r StudentController) PUT(c echo.Context) error {

	studententity := &entity.Student{}

	if err := c.Bind(studententity); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, studententity)
}

func (r StudentController) DELETE(c echo.Context) error {

	return c.JSON(http.StatusOK, nil)
}
