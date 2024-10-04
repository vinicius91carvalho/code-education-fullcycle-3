package graph

import "github.com/vinicius91carvalho/code-education-graphql/internal/database"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	CategoryDB *database.Category 
}
