package db

import (
	"context"
	"database/sql"
	"log"

	entity "github.com/RomeroGabriel/go-graphQL/internal/entity"
	"github.com/google/uuid"
)

type CategoryDB struct {
	db *sql.DB
}

func NewCategory(db *sql.DB) *CategoryDB {
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS Categories 
		(id TEXT NOT NULL, name TEXT NOT NULL, description TEXT NOT NULL);
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
	return &CategoryDB{db: db}
}

func (c *CategoryDB) Create(name string, description string) (entity.Category, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO Categories (id, name, description) VALUES ($1, $2, $3)",
		id, name, description)

	if err != nil {
		return entity.Category{}, err
	}
	return entity.Category{Id: id, Name: name, Description: description}, nil
}

func (c *CategoryDB) FindAll() ([]entity.Category, error) {
	rows, err := c.db.Query("SELECT id, name, description FROM Categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	categories := []entity.Category{}
	for rows.Next() {
		var id, name, description string
		if err := rows.Scan(&id, &name, &description); err != nil {
			return nil, err
		}
		categories = append(categories, entity.Category{Id: id, Name: name, Description: description})
	}
	return categories, nil
}
