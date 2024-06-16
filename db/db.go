package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Db() *gorm.DB {
	Db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return Db
}
