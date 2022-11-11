package repository

import (
	"adla.com/task10.11.2022/entity"
	"gorm.io/gorm"
)

type StudentRepository struct {
	DB *gorm.DB
}

func (r StudentRepository) GetAll() []entity.Student {

	var student []entity.Student
	_ = r.DB.Find(&student)
	return student
}
