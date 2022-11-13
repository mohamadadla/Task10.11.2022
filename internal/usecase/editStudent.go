package usecase

import (
	"adla.com/task10.11.2022/internal/adapter/db/repository"
	"adla.com/task10.11.2022/internal/interface/api/payload"
)

type EditStudent struct {
	RepositoryStudent repository.StudentRepository
}

func (e EditStudent) Handle(studentid *payload.EditStudent) error {
	if err := e.RepositoryStudent.Put(studentid); err != nil {
		return err
	}
	return nil
}
