package database

import (
	"github.com/dgraph-io/badger/v3"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

// NewSQLDatabase create a gorm database object.
// kind: type of database, like "sqlite", "mysql", "postgresql", etc.
// dsn: dsn with user/password and all necessary for connect in database.
func NewSQLDatabase(kind, dsn string) (db *gorm.DB) {
	switch kind {
	case "sqlite":
		db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect in sqlite database")
		}
		return db
	case "mysql":
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect in mysql database")
		}
		return db
	case "postgresql":
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
