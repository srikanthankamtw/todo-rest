package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func Init() *gorm.DB {
	db, err := gorm.Open("postgres",
		"host=127.0.0.1 port=5432 user=postgres password=postgres dbname=todo sslmode=disable")

	if err != nil {
		panic(err.Error())
	}

	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
