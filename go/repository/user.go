package repository

import (
	"github.com/Kento-Ishizaki/go-next-template/entity"
	"github.com/Kento-Ishizaki/go-next-template/repository/internal/mock"
	"github.com/Kento-Ishizaki/go-next-template/repository/internal/mysql"
)

type UserRepository interface {
	GetUser(id int) (*entity.User, error)
	CreateUser(name string) (*entity.User, error)
}

func (m *MysqlRepositoryFactory) UserRepo() UserRepository {
	return &mysql.MysqlUserRepo{DB: m.db}
}

func (m *MockRepositoryFactory) UserRepo() UserRepository {
	return &mock.MockUserRepo{}
}
