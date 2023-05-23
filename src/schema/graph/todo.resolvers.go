package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"fmt"

	"github.com/irdaislakhuafa/go-graphql-upload-files/src/schema/graph/converts"
	"github.com/irdaislakhuafa/go-graphql-upload-files/src/schema/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo, err := converts.NewTodoToTodo(input)
	if err != nil {
		return nil, err
	}

	todo, err = r.Todo.Save(ctx, todo)
	if err != nil {
		return nil, err
	}

	result := model.Todo{
		ID:   todo.ID,
		Text: todo.Text,
		Done: todo.Done,
		User: &model.User{
			ID:   "",
			Name: "",
		},
	}
	return &result, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented: Todos - todos"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
