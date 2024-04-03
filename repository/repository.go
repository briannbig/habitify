package repository

import (
	"github.com/briannbig/habitify/db"
	"github.com/briannbig/habitify/model"
)

var (
	database = db.New().Connection
)

type (
	Repository[M model.Model] interface {
		Create(model M) (M, error)
		GetAll() []M
		Update(id uint, model M) (M, error)
		Delete(id uint) error
	}

	UserRepository struct {
	}

	habitRepository struct {
	}
)

func NewUserRepository() Repository[model.User] {
	return UserRepository{}
}

func NewHabitRepository() Repository[model.Habit] {
	return habitRepository{}
}

func (u UserRepository) Create(user model.User) (model.User, error) {
	id := database.Create(
		&model.User{
			Email:    user.Email,
			Password: user.Password,
		})
	var savedUser model.User
	database.First(&savedUser, id)
	return savedUser, nil
}

func (u UserRepository) Delete(id uint) error {
	database.Delete(&model.User{}, id)
	return nil
}

func (u UserRepository) GetAll() []model.User {
	var users []model.User
	database.Find(&users)
	return users
}

func (u UserRepository) Update(id uint, user model.User) (model.User, error) {
	var user_ model.User
	database.First(&model.User{}, id)

	if user.Email != "" {
		user_.Email = user.Email
	}

	if user.Password != "" {
		user_.Password = user.Password
	}

	database.Save(user_)

	database.First(user, id)

	return user, nil

}

func (u UserRepository) FindUserByEmail(email string) (model.User, error) {
	var user model.User
	database.Where("email = ?", email).First(&user)

	return user, nil
}

func (h habitRepository) Create(habit model.Habit) (model.Habit, error) {
	id := database.Create(&habit)
	var savedHabit model.Habit
	database.First(&savedHabit, id)
	return savedHabit, nil
}

func (h habitRepository) Delete(id uint) error {
	database.Delete(&model.Habit{}, id)
	return nil
}

func (h habitRepository) GetAll() []model.Habit {
	var habits []model.Habit
	database.Find(&habits)
	return habits
}

func (h habitRepository) Update(id uint, habit model.Habit) (model.Habit, error) {
	var habit_ model.Habit
	database.First(&habit_, id)

	if habit.Name != "" {
		habit_.Name = habit.Name
	}

	if habit_.Description != "" {
		habit_.Description = habit.Description
	}

	if habit.CompletionStatus != "" {
		habit_.CompletionStatus = habit.CompletionStatus
	}

	database.Save(&habit_)

	return habit_, nil
}
