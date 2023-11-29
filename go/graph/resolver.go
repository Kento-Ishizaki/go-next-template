package graph

import "github.com/Kento-Ishizaki/go-next-template/repository"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// type Usecase interface {
// 	CreateUser(user *model.User)
// }

// type UserUsecase struct {
// 	ur repository.UserRepository
// }

// type UserRepository interface {
// 	CreateUser(name string) (string, error)
// }

// type MockRepo struct{}

// func NewMockUserRepo() Repository {
// 	return &MockRepo{}
// }

// func (*MockRepo) CreateUser(name string) (string, error) {
// 	return "1", nil
// }

// type MysqlRepo struct {
// 	DB *gorm.DB
// }

// func NewMysqlUserRepo(db *gorm.DB) Repository {
// 	return &MysqlRepo{
// 		DB: db,
// 	}
// }

// func (r *MysqlRepo) CreateUser(name string) (string, error) {
// 	user := &models.User{Name: name}
// 	res := r.DB.Create(user)
// 	if err := res.Error; err != nil {
// 		return "", err
// 	}

// 	return strconv.Itoa(user.ID), nil
// }

type Resolver struct {
	// interfaceを入れる
	Repo repository.RepositoryFactory
}
