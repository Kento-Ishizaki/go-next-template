package mysql

import (
	"github.com/Kento-Ishizaki/go-next-template/entity"
	"gorm.io/gorm"
)

type MysqlTodoRepo struct {
	DB *gorm.DB
}

func (r *MysqlTodoRepo) GetTodo(id int) (*entity.Todo, error) {
	t := new(entity.Todo)
	if err := r.DB.First(t, id).Error; err != nil {
		return nil, err
	}

	return t, nil
}

func (r *MysqlTodoRepo) CreateTodo(text string, userId int) (*entity.Todo, error) {
	t := &entity.Todo{Text: text, UserID: userId}
	if err := r.DB.Create(t).Error; err != nil {
		return nil, err
	}

	return t, nil
}
