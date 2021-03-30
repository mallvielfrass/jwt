package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func Init(name string) Database {
	db, err := gorm.Open(sqlite.Open(name), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return Database{
		db: db,
	}
}
