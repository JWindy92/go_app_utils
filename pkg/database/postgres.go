package database

import (
	"fmt"
	"log"

	"github.com/JWindy92/go_app_utils/pkg/logging"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresImpl struct {
	GormDB
}

func (impl *PostgresImpl) ConnectDB() *gorm.DB {
	// Example: replace with real values or use environment variables
	host := "localhost"
	port := 5555
	user := "postgres"
	password := "dbpass"
	dbname := "postgres"

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	log.Printf("Connecting to Postgres DB: %s\n", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logging.GormLog,
	})
	if err != nil {
		panic(fmt.Sprintf("failed to connect to Postgres: %v", err))
	}

	impl.DB = db
	return db
}
