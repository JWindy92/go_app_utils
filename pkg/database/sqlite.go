package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLiteImpl struct {
	GormDB
}

func (impl *SQLiteImpl) ConnectDB(dbPath string) *gorm.DB {
	log.Printf("Connecting to SQLite DB %s\n", dbPath)
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	impl.DB = db
	return db
}
