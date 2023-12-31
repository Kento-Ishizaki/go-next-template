package core

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewDB is a function to connect to the database.
func NewDB(cfg *Config) (*gorm.DB, error) {
	dns := cfg.DB.User + ":" + cfg.DB.Password + "@tcp(" + cfg.DB.Host + ":" + cfg.DB.Port + ")/" + cfg.DB.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	return gorm.Open(mysql.New(mysql.Config{
		DSN:                       dns,   // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})
}
