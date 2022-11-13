package usecase

import (
	"adla.com/task10.11.2022/internal/adapter/db/entity"
	"adla.com/task10.11.2022/internal/adapter/db/repository"
)

type FindStudentById struct {
	RepositoryStudent repository.StudentRepository
}

func (f FindStudentById) Handle(studentid int) (*entity.Student, error) {

	student, err := f.RepositoryStudent.Get(studentid)

	if err != nil {
		return nil, err
	}

	return student, nil
}
