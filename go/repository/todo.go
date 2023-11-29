package repository

import (
	"github.com/Kento-Ishizaki/go-next-template/entity"
	"github.com/Kento-Ishizaki/go-next-template/repository/internal/mock"
	"github.com/Kento-Ishizaki/go-next-template/repository/internal/mysql"
)

type TodoRepository interface {
	GetTodo(id int) (*entity.Todo, error)
}

func (m *MockRepositoryFactory) TodoRepo() TodoRepository {
	return &mock.MockTodoRepo{}
}

func (m *MysqlRepositoryFactory) TodoRepo() TodoRepository {
	return &mysql.MysqlTodoRepo{DB: m.db}
}
