package usecase

import (
	"adla.com/task10.11.2022/internal/adapter/db/entity"
	"adla.com/task10.11.2022/internal/adapter/db/repository"
)

type GetAllStudents struct {
	RepositoryStudent repository.StudentRepository
}

func (g GetAllStudents) Handle() ([]*entity.Student, error) {

	students, err := g.RepositoryStudent.GetAll()

	if err != nil {
		return nil, err
	}
	return students, nil
}
