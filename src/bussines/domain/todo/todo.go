package todo

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/irdaislakhuafa/go-graphql-upload-files/src/config"
	"github.com/irdaislakhuafa/go-graphql-upload-files/src/schema/entity"
	"github.com/irdaislakhuafa/go-graphql-upload-files/src/utils/files"
)

type Interface interface {
	Save(ctx context.Context, todo entity.Todo) (entity.Todo, error)
	Get(ctx context.Context, todo entity.Todo) (entity.Todo, error)
	GetAll(ctx context.Context, todo entity.Todo) ([]entity.Todo, []byte, error)
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
			log.Println(err)
			return entity.Todo{}, err
		}
	}

	fileID, err := files.SaveFile(ctx, t.config.Database.DirectoryFiles, uuid.NewString()+"-"+todo.File.Filename, todo.File, 0755)
	if err != nil {
		log.Println(err)
		return entity.Todo{}, err
	}
	todo.FileID = fileID

	todos, tempBytes, err := t.GetAll(ctx, entity.Todo{})
	if err != nil {
		log.Println(err)
		return entity.Todo{}, err
	}

	todos = append(todos, todo)
	todoBytes, err := json.Marshal(todos)
	if err != nil {
		log.Println(err)
		return entity.Todo{}, err
	}

	if err := os.WriteFile(t.store, todoBytes, 0644); err != nil {
		log.Println(err)
		err := os.WriteFile(t.store, tempBytes, 0644)
		if err != nil {
			log.Println(err)
			return entity.Todo{}, err
		}
		return entity.Todo{}, err
	}

	return todo, nil
}

func (t *todo) Get(ctx context.Context, todo entity.Todo) (entity.Todo, error) {
	todos, _, err := t.GetAll(ctx, todo)
	if err != nil {
		return entity.Todo{}, err
	}

	for _, v := range todos {
		if todo.ID == v.ID {
			return v, nil
		}
	}

	return entity.Todo{}, fmt.Errorf("not found")
}

func (t *todo) GetAll(ctx context.Context, todo entity.Todo) ([]entity.Todo, []byte, error) {
	tempBytes, err := os.ReadFile(t.store)
	if err != nil {
		log.Println(err)
		return nil, tempBytes, err
	}

	todos := []entity.Todo{}
	if err := json.Unmarshal(tempBytes, &todos); err != nil {
		log.Println(err)
		return nil, tempBytes, err
	}

	return todos, nil, nil
}

func (t *todo) Update(ctx context.Context, todo entity.Todo) (entity.Todo, error) {
	todos, tempBytes, err := t.GetAll(ctx, todo)
	if err != nil {
		return entity.Todo{}, err
	}

	for i := range todos {
		if todos[i].ID == todo.ID {
			todos[i].UserID = todo.UserID
			todos[i].Text = todo.Text
			todos[i].FileID = todo.FileID
			todos[i].Done = todo.Done

			bytesStream, err := json.Marshal(todos)
			if err != nil {
				return entity.Todo{}, err
			}

			if err := os.WriteFile(t.store, bytesStream, 0644); err != nil {
				if err := os.WriteFile(t.store, tempBytes, 0644); err != nil {
					return entity.Todo{}, err
				}
				return entity.Todo{}, err
			}

			return todos[i], nil
		}
	}

	return entity.Todo{}, fmt.Errorf("no rows updated")

}

func (t *todo) Delete(ctx context.Context, todo entity.Todo) (entity.Todo, error) {
	todos, tempBytes, err := t.GetAll(ctx, todo)
	if err != nil {
		return entity.Todo{}, err
	}

	for i := range todos {
		if todos[i].ID == todo.ID {
			todo = todos[i]
			todos = append(todos[:i], todos[i+1:]...)
			break
		}
	}

	bytesStream, err := json.Marshal(todos)
	if err != nil {
		return entity.Todo{}, err
	}

	if err := os.WriteFile(t.store, bytesStream, 0644); err != nil {
		if err := os.WriteFile(t.store, tempBytes, 0644); err != nil {
			return entity.Todo{}, err
		}
		return entity.Todo{}, err
	}

	return todo, nil
}
