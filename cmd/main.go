package main

import (
	"log/slog"

	"github.com/JWindy92/go_app_utils/pkg/database"
	"github.com/JWindy92/go_app_utils/pkg/logging"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// ID           uint         `gorm:"primaryKey" json:"id"`
	Name         string `json:"name,omitempty"`
	Email        string `gorm:"unique" json:"email"`
	PasswordHash string `json:"password_hash"`
}

func main() {
	logging.InitLogger()
	slog.Info("Logger initialized")
	sqlite := database.SQLiteImpl{}
	sqlite.ConnectDB("test.db")
	sqlite.RunAutoMigrate(&User{})
}
