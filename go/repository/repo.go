package repository

import (
	"github.com/Kento-Ishizaki/go-next-template/core"
	"gorm.io/gorm"
)

// RepositoryFactory はリポジトリのインターフェースを定義します。
type RepositoryFactory interface {
	UserRepo() UserRepository
	TodoRepo() TodoRepository
}

// MockRepositoryFactory はモックバージョンのリポジトリを作成します。
type MockRepositoryFactory struct{}

// MysqlRepositoryFactory はMySQLバージョンのリポジトリを作成します。
type MysqlRepositoryFactory struct {
	db *gorm.DB
}

// NewRepositoryFactory はリポジトリのファクトリを作成します。
func NewRepositoryFactory(cfg *core.Config, db *gorm.DB) RepositoryFactory {
	if cfg.UseMock {
		return &MockRepositoryFactory{}
	}

	return &MysqlRepositoryFactory{
		db: db,
	}
}
