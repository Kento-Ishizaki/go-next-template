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
