package todo

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/irdaislakhuafa/go-graphql-upload-files/src/config"
	"github.com/irdaislakhuafa/go-graphql-upload-files/src/schema/entity"
)

type Interface interface {
	Save(ctx context.Context, todo entity.Todo) (entity.Todo, error)
	Get(ctx context.Context, todo entity.Todo) (entity.Todo, error)
	GetAll(ctx context.Context, todo entity.Todo) ([]entity.Todo, error)
	Update(ctx context.Context, todo entity.Todo) (entity.Todo, error)
	Delete(ctx context.Context, todo entity.Todo) (entity.Todo, error)
}

type todo struct {
	store  string
	config config.Config
}

func Init(ctx context.Context, config config.Config) Interface {
	return &todo{
		store:  fmt.Sprintf("%s/%s", config.Database.Directory, "todo.json"),
		config: config,
	}
}

func (t *todo) Save(ctx context.Context, todo entity.Todo) (entity.Todo, error) {
	todo.ID = uuid.NewString()
	_, err := os.Stat(t.store)
	if os.IsNotExist(err) {
		if err := os.WriteFile(t.store, []byte("[]"), 0644); err != nil {
			return entity.Todo{}, err
		}
	}

	tempBytes, err := os.ReadFile(t.store)
	if err != nil {
		return entity.Todo{}, err
	}

	todos := []entity.Todo{}
	if err := json.Unmarshal(tempBytes, &todos); err != nil {
		return entity.Todo{}, err
	}

	todos = append(todos, todo)
	todoBytes, err := json.Marshal(todos)
	if err != nil {
		return entity.Todo{}, err
	}

	if err := os.WriteFile(t.store, todoBytes, 0644); err != nil {
		err := os.WriteFile(t.store, tempBytes, 0644)
		if err != nil {
			return entity.Todo{}, err
		}
		return entity.Todo{}, err
	}

	return todo, nil
}

func (t *todo) Get(ctx context.Context, todo entity.Todo) (entity.Todo, error) {
	panic("not implemented") // TODO: Implement
}

func (t *todo) GetAll(ctx context.Context, todo entity.Todo) ([]entity.Todo, error) {
	panic("not implemented") // TODO: Implement
}

func (t *todo) Update(ctx context.Context, todo entity.Todo) (entity.Todo, error) {
	panic("not implemented") // TODO: Implement
}

func (t *todo) Delete(ctx context.Context, todo entity.Todo) (entity.Todo, error) {
	panic("not implemented") // TODO: Implement
}
