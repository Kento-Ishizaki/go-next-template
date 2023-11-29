package mysql

import (
	"github.com/Kento-Ishizaki/go-next-template/entity"
	"gorm.io/gorm"
)

type MysqlUserRepo struct {
	DB *gorm.DB
}

func (r *MysqlUserRepo) GetUser(id int) (*entity.User, error) {
	u := new(entity.User)
	if err := r.DB.First(u, id).Error; err != nil {
		return nil, err
	}

	return u, nil
}
