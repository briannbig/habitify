package db

import (
	"github.com/briannbig/habitify/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Db struct {
	Connection *gorm.DB
}

func New() Db {

	db, err := gorm.Open(sqlite.Open("habitify.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to the database")
	}

	db.AutoMigrate(&model.User{}, &model.Habit{})

	return Db{
		Connection: db,
	}

}
