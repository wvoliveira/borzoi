package database

import (
	"log"
	"os"
	"time"

	"github.com/dgraph-io/badger/v3"
	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewSQLDatabase create a gorm database object.
// kind: type of database, like "sqlite", "mysql", "postgresql", etc.
// dsn: dsn with user/password and all necessary for connect in database.
func NewSQLDatabase(kind, dsn string) (db *gorm.DB) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)

	cfg := gorm.Config{Logger: newLogger}

	switch kind {
	case "sqlite":
		db, err := gorm.Open(sqlite.Open(dsn), &cfg)
		if err != nil {
			panic("failed to connect in sqlite database")
		}
		return db

	case "mysql":
		db, err := gorm.Open(mysql.Open(dsn), &cfg)
		if err != nil {
			panic("failed to connect in mysql database")
		}
		return db

	case "postgresql":
		db, err := gorm.Open(postgres.Open(dsn), &cfg)
		if err != nil {
			panic("failed to connect in postgresql database")
		}
		return db

	default:
		log.Fatal("this type of database is not supported")
	}
	return
}

// NewNoSQLDatabase create a noSQL database client (only badger yet).
// kind: type of database, like "badger", "redis", "bolt", etc.
// dsn: dsn with user/password and all necessary for connect in database.
func NewNoSQLDatabase(kind, dsn string) (db *badger.DB) {
	switch kind {
	case "badger":
		db, err := badger.Open(badger.DefaultOptions(dsn))
		if err != nil {
			log.Fatalf("error to connect in badger database: %s", err)
		}
		return db

	default:
		log.Fatal("this type of database is not supported")
	}
	return
}
