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
