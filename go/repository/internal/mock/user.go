package mock

import (
	"github.com/Kento-Ishizaki/go-next-template/entity"
)

type MockUserRepo struct{}

func (*MockUserRepo) GetUser(id int) (*entity.User, error) {
	return &entity.User{
		ID:   id,
		Name: "test",
	}, nil
}

func (*MockUserRepo) CreateUser(name string) (*entity.User, error) {
	return &entity.User{
		ID:   1,
		Name: name,
	}, nil
}
