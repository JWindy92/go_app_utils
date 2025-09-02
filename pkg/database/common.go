package database

import (
	"gorm.io/gorm"
)

type GormDB struct {
	DB *gorm.DB
}

func (db *GormDB) RunAutoMigrate(models ...interface{}) error {
	return db.DB.AutoMigrate(models)
}
