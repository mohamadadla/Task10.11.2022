package usecase

import (
	"adla.com/task10.11.2022/internal/adapter/db/repository"
	"adla.com/task10.11.2022/internal/interface/api/payload"
)

type CreateStudent struct {
	RepositoryStudent repository.StudentRepository
}

func (c CreateStudent) Handle(student *payload.NewStudnet) error {
	if err := c.RepositoryStudent.Create(student); err != nil {
		return err
	}
	return nil
}
