package db

import (
	"database/sql"

	entity "github.com/RomeroGabriel/go-graphQL/internal/entity"
	"github.com/google/uuid"
)

type CategoryDB struct {
	db *sql.DB
}

func NewCategory(db *sql.DB) *CategoryDB {
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
