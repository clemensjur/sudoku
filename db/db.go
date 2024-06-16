package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Db(path string) *gorm.DB {
	Db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return Db
}
