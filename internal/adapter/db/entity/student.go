package entity

type Student struct {
	ID        int
	Firstname string
	Lastname  string
}

func (s Student) TableName() string {
	return "students"
}
