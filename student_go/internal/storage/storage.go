package storage

import "github.com/kapilrohilla/codebase/internal/types"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudent(offset int64, limit int64) (interface{}, error)
	GetStudentById(id int64) (types.Student, error)
	UpdateStudentById(id int64) (int64, error)
	DeleteStudentById(id int64) (int64, error)
}
