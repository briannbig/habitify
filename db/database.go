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

	initSystemUser(db)

	return Db{
		Connection: db,
	}

}

func initSystemUser(db *gorm.DB) {
	var user model.User
	db.Where("email = ?", "system").First(&user)

	if user.Email == "" {
		db.Create(&model.User{Email: "system", Password: "system", Role: model.UserRoleAdmin})
	}
}
