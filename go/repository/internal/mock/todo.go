package mock

import (
	"strconv"

	"github.com/Kento-Ishizaki/go-next-template/entity"
)

type MockTodoRepo struct{}

func (*MockTodoRepo) GetTodo(id int) (*entity.Todo, error) {
	return &entity.Todo{
		ID:   id,
		Text: strconv.Itoa(id),
	}, nil
}

func (*MockTodoRepo) CreateTodo(text string, userId int) (*entity.Todo, error) {
	return &entity.Todo{
		ID:   userId,
		Text: text,
	}, nil
}
