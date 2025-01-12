package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/kapilrohilla/codebase/internal/config"
	"github.com/kapilrohilla/codebase/internal/types"
	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	DB *sql.DB
}

// CreateStudent implements storage.Storage.
func (s *Sqlite) CreateStudent(name string, email string, age int) (int64, error) {
	stmt, err := s.DB.Prepare("INSERT INTO students (name,email, age) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(name, email, age)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *Sqlite) GetStudent(offset int64, limit int64) (interface{}, error) {
	result, err := s.DB.Query("SELECT * from students s ")

	if err != nil {
		return nil, err
	}

	defer result.Close()
	var studs []types.Student

	for result.Next() {
		var stud types.Student
		err := result.Scan(&stud.Id, &stud.Name, &stud.Age, &stud.Email, &stud.CreatedAT)
		if err != nil {
			fmt.Println(err)
		}
		studs = append(studs, stud)
	}

	return studs, nil
}
func New(cfg config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.Storage)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS students (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(50) NOT NULL,
			age TINYINT DEFAULT NULL, 
			email TEXT NOT NULL,
			createdAT DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`)

	if err != nil {
		return nil, err
	}
	return &Sqlite{DB: db}, err
}
