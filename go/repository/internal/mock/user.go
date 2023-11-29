package mock

import (
	"github.com/Kento-Ishizaki/go-next-template/entity"
)

type MockUserRepo struct{}

func (*MockUserRepo) GetUser(id int) (*entity.User, error) {
	return &entity.User{
		ID:   id,
		Name: "tes",
	}, nil
}
