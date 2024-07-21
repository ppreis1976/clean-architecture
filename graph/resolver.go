package graph

import "clean-architecture/internal/repository"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your graphql, add any dependencies you require here.

type Resolver struct {
	OrderRepository repository.OrderRepository
}
