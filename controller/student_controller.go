package controller

import (
	"net/http"

	"adla.com/task10.11.2022/entity"
	"github.com/labstack/echo/v4"
)

type StrudentController struct {
}

func (receiver StrudentController) GetAll(c echo.Context) error {

	studententities := []entity.Student{}
	return c.JSON(http.StatusOK, studententities)
}

func (receiver StrudentController) Get(c echo.Context) error {

	studententity := &entity.Student{}
	return c.JSON(http.StatusOK, studententity)
}

func (receiver StrudentController) Create(c echo.Context) error {

	studententity := &entity.Student{}

	if err := c.Bind(studententity); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, studententity)
}

func (receiver StrudentController) PUT(c echo.Context) error {

	studententity := &entity.Student{}

	if err := c.Bind(studententity); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, studententity)
}

func (receiver StrudentController) DELETE(c echo.Context) error {

	return c.JSON(http.StatusOK, nil)
}
