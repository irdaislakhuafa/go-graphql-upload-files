package converts

import (
	"github.com/irdaislakhuafa/go-graphql-upload-files/src/schema/entity"
	"github.com/irdaislakhuafa/go-graphql-upload-files/src/schema/graph/model"
)

func NewTodoToTodo(input model.NewTodo) (entity.Todo, error) {
	todo := entity.Todo{
		ID:     "",
		Text:   input.Text,
		Done:   false,
		UserID: input.UserID,
	}

	if input.File != nil {
		todo.File = *input.File
	}

	return todo, nil
}
