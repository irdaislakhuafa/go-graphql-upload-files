package graph

import "github.com/irdaislakhuafa/go-graphql-upload-files/src/bussines/usecase/todo"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Todo todo.Interface
}
