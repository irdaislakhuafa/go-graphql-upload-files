package todo

import (
	"context"

	domainTodo "github.com/irdaislakhuafa/go-graphql-upload-files/src/bussines/domain/todo"
	"github.com/irdaislakhuafa/go-graphql-upload-files/src/schema/entity"
)

type Interface interface {
	Save(ctx context.Context, todo entity.Todo) (entity.Todo, error)
}

type todo struct {
	domainTodo domainTodo.Interface
}

func Init(ctx context.Context, domainTodo domainTodo.Interface) Interface {
	return &todo{
		domainTodo: domainTodo,
	}
}

func (t *todo) Save(ctx context.Context, todo entity.Todo) (entity.Todo, error) {
	todo, err := t.domainTodo.Save(ctx, todo)
	if err != nil {
		return entity.Todo{}, err
	}

	return todo, nil
}
