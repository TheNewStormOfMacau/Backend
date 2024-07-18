package core

import (
	"backend/global"
	"backend/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

func Gorm() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func Init() {
	global.G_DB = Gorm()
	err := global.G_DB.AutoMigrate(&model.User{}, &model.Record{})
	if err != nil {
		os.Exit(0)
	}
	Gin()
}
