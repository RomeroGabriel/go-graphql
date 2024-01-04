package db

import (
	"context"
	"database/sql"
	"log"

	entity "github.com/RomeroGabriel/go-graphQL/internal/entity"
	"github.com/google/uuid"
)

type CourseDB struct {
	db *sql.DB
}

func NewCourseDB(db *sql.DB) *CourseDB {
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS Courses 
		(id TEXT NOT NULL, name TEXT NOT NULL, description TEXT NOT NULL, categoryId TEXT NOT NULL);
	`
	ctx := context.Background()
	conn, err := db.Conn(ctx)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return nil
	}
	defer conn.Close()
	_, err = conn.ExecContext(ctx, sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return nil
	}
	return &CourseDB{db: db}
}

func (c *CourseDB) Create(name, description, categoryId string) (entity.Course, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO Courses (id, name, description, categoryID) VALUES ($1, $2, $3, $4)",
		id, name, description, categoryId)

	if err != nil {
		return entity.Course{}, err
	}
	return entity.Course{Id: id, Name: name, Description: description, CategoryId: categoryId}, nil
}

func (c *CourseDB) FindAll() ([]entity.Course, error) {
	rows, err := c.db.Query("SELECT id, name, description, categoryID FROM Courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	courses := []entity.Course{}
	for rows.Next() {
		var id, name, description, categoryId string
		if err := rows.Scan(&id, &name, &description, &categoryId); err != nil {
			return nil, err
		}
		courses = append(courses, entity.Course{Id: id, Name: name, Description: description, CategoryId: categoryId})
	}
	return courses, nil
}

func (c *CourseDB) FindByCategoryId(categoryId string) ([]entity.Course, error) {
	rows, err := c.db.Query("SELECT id, name, description, categoryID FROM Courses WHERE categoryId = $1", categoryId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	courses := []entity.Course{}
	for rows.Next() {
		var id, name, description, categoryId string
		if err := rows.Scan(&id, &name, &description, &categoryId); err != nil {
			return nil, err
		}
		courses = append(courses, entity.Course{Id: id, Name: name, Description: description, CategoryId: categoryId})
	}
	return courses, nil
}
