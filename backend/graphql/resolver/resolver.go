package resolver

import (
	"gorm.io/gorm"
)

// Resolver is the main struct for the GraphQL resolver
type Resolver struct {
	DB *gorm.DB // Menyimpan instance GORM DB untuk digunakan di resolver
}

// NewResolver initializes a new Resolver
func NewResolver(db *gorm.DB) *Resolver {
	return &Resolver{
		DB: db,
	}
}

