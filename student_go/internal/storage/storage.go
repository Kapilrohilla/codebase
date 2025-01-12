package storage

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudent(offset int64, limit int64) (interface{}, error)
}
