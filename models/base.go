package models

import (
	"github.com/go-redis/redis/v7"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var REDIS *redis.Client

func init() {
	db, err := gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	DB = db
	err = DB.AutoMigrate(&User{})
	if err != nil {
		panic(err.Error())
	}

	redis := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	REDIS = redis
}
