package api

import (
	"adla.com/task10.11.2022/internal/interface/api/payload"
	"adla.com/task10.11.2022/internal/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type StudentController struct {
	GetAllStudents usecase.GetAllStudents
	FindById       usecase.FindStudentById
	CreateStudent  usecase.CreateStudent
	DeleteStudent  usecase.DeleteStudent
	EditStudent    usecase.EditStudent
}

func (s StudentController) GetAll(c echo.Context) error {

	students, err := s.GetAllStudents.Handle()

	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	return c.JSON(http.StatusOK, students)
}

func (s StudentController) Get(c echo.Context) error {

	studentById := &payload.StudentByID{}
	if err := c.Bind(studentById); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	student, err := s.FindById.Handle(studentById.Id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "x")
	}
	return c.JSON(http.StatusOK, student)
}

func (s StudentController) Create(c echo.Context) error {
	newStudent := &payload.NewStudnet{}

	if err := c.Bind(newStudent); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := s.CreateStudent.Handle(newStudent); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusCreated, nil)
}

func (s StudentController) Put(c echo.Context) error {
	editStudent := payload.EditStudent{}

	if err := c.Bind(&editStudent); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := s.EditStudent.Handle(&editStudent); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, nil)
}

func (s StudentController) Delete(c echo.Context) error {
	studentId := &payload.DeleteStudent{}

	if err := c.Bind(studentId); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	err := s.DeleteStudent.Handle(studentId.Id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, nil)
}
