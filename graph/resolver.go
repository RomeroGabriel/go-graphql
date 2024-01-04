package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import "github.com/RomeroGabriel/go-graphQL/internal/db"

type Resolver struct {
	CategoryDB *db.CategoryDB
	CourseDB   *db.CourseDB
}
