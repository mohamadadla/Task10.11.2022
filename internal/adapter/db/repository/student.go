package repository

import (
	"adla.com/task10.11.2022/internal/adapter/db/entity"
	"adla.com/task10.11.2022/internal/domain/repository"
	"adla.com/task10.11.2022/internal/interface/api/payload"
	"gorm.io/gorm"
)

type StudentRepository struct {
	DB *gorm.DB
}

func ProvideStudentRepository(db *gorm.DB) repository.StudentRepository {
	return StudentRepository{
		DB: db,
	}
}

func (s StudentRepository) GetAll() ([]*entity.Student, error) {

	var students []*entity.Student
	if db := s.DB.Find(&students); db.Error != nil {
		return nil, db.Error
	}
	return students, nil
}

func (s StudentRepository) Get(id int) (*entity.Student, error) {
	var student *entity.Student
	if db := s.DB.Find(&student, "id =?", id); db.Error != nil {
		return nil, db.Error
	}
	return student, nil
}

func (s StudentRepository) Create(student *payload.NewStudnet) error {
	newStudent := &entity.Student{
		Firstname: student.FirstName,
		Lastname:  student.LastName,
	}
	if db := s.DB.Create(newStudent); db.Error != nil {
		return db.Error
	}
	return nil
}

func (s StudentRepository) Put(student *payload.EditStudent) error {
	var editStudent = &entity.Student{
		ID: student.Id,
	}
	result := s.DB.First(&editStudent)

	if result.RowsAffected > 0 {
		editStudent.Firstname = student.FirstName
		editStudent.Lastname = student.LastName
		s.DB.Save(&editStudent)
	}
	return nil
}

func (s StudentRepository) Delete(id int) error {
	if db := s.DB.Delete(&entity.Student{}, "id =?", id); db.Error != nil {
		return db.Error
	}
	return nil
}
