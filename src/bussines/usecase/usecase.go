package usecase

import (
	"context"

	"github.com/irdaislakhuafa/go-graphql-upload-files/src/bussines/domain"
	"github.com/irdaislakhuafa/go-graphql-upload-files/src/bussines/usecase/todo"
	"github.com/irdaislakhuafa/go-graphql-upload-files/src/config"
)

type Usecase struct {
	Todo todo.Interface
}

func Init(ctx context.Context, config config.Config, domain domain.Domain) *Usecase {
	return &Usecase{
		Todo: todo.Init(ctx, domain.Todo),
	}
}
