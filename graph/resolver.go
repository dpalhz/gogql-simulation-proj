package graph

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

// Di sini Anda dapat menambahkan fungsi atau metode lainnya yang akan digunakan di resolver
