package usecase

import "adla.com/task10.11.2022/internal/adapter/db/repository"

type DeleteStudent struct {
	RepositoryStudent repository.StudentRepository
}

func (d DeleteStudent) Handle(studentId int) error {
	err := d.RepositoryStudent.Delete(studentId)

	if err != nil {
		return err
	}
	return nil
}
