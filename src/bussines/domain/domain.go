package domain

import (
	"context"

	"github.com/irdaislakhuafa/go-graphql-upload-files/src/bussines/domain/todo"
	"github.com/irdaislakhuafa/go-graphql-upload-files/src/config"
)

type Domain struct {
	Todo todo.Interface
}

func Init(ctx context.Context, config config.Config) *Domain {
	return &Domain{
		Todo: todo.Init(ctx, config),
	}
}
