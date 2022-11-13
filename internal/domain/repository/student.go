package repository

import (
	"adla.com/task10.11.2022/internal/adapter/db/entity"
	"adla.com/task10.11.2022/internal/interface/api/payload"
)

type StudentRepository interface {
	GetAll() ([]*entity.Student, error)
	Create(student *payload.NewStudnet) error
	Get(id int) (*entity.Student, error)
	Put(student *payload.EditStudent) error
	Delete(id int) error
}
