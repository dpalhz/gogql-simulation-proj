package resolver

import "notes/backend/services"

// Resolver holds the grouped services
type Resolver struct {
	Services *services.Services
}

// NewResolver initializes a Resolver with the grouped services.
func NewResolver(services *services.Services) *Resolver {
	return &Resolver{
		Services: services,
	}
}
