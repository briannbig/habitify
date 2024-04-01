package repository

import (
	"github.com/briannbig/habitify/db"
	"github.com/briannbig/habitify/model"
)

var (
	database = db.New()
)

type (
	Repository[M model.Model] interface {
		Create(model M) (M, error)
		GetAll() []M
		Update(id uint, model M) (M, error)
		Delete(id uint) error
	}

	userRepository struct {
		db db.Db
	}

	habitRepository struct {
		db db.Db
	}
)

func NewUserRepository() Repository[model.User] {
	return userRepository{db: database}
}

func NewHabitRepository() Repository[model.Habit] {
	return habitRepository{db: database}
}

func (u userRepository) Create(model model.User) (model.User, error) {
	panic("unimplemented")
}

func (u userRepository) Delete(id uint) error {
	panic("unimplemented")
}

func (u userRepository) GetAll() []model.User {
	panic("unimplemented")
}

func (u userRepository) Update(id uint, model model.User) (model.User, error) {
	panic("unimplemented")
}

func (h habitRepository) Create(model model.Habit) (model.Habit, error) {
	panic("unimplemented")
}

func (h habitRepository) Delete(id uint) error {
	panic("unimplemented")
}

func (h habitRepository) GetAll() []model.Habit {
	panic("unimplemented")
}

func (h habitRepository) Update(id uint, model model.Habit) (model.Habit, error) {
	panic("unimplemented")
}
