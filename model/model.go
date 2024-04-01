package model

import (
	"encoding/json"
	"log"

	"gorm.io/gorm"
)

type HabitCompletionStatus string

const (
	Started   HabitCompletionStatus = "Started"
	Completed HabitCompletionStatus = "Completed"
	Done      HabitCompletionStatus = "Done"
	Missed    HabitCompletionStatus = "Missed"
	NotDone   HabitCompletionStatus = "Not done"
)

type (
	Model interface {
		User | Habit
	}

	User struct {
		gorm.Model
		Email    string  `json:"email"`
		Password string  `json:"-"`
		Habits   []Habit `gorm:"foreignKey:UserID"`
	}

	Habit struct {
		gorm.Model
		UserId           uint                  `json:"userId"`
		Name             string                `json:"name"`
		Description      string                `json:"description,omitempty"`
		CompletionStatus HabitCompletionStatus `json:"habit-completion-status"`
	}
)

func (user User) String() string {
	out, err := json.MarshalIndent(user, "", "   ")
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}

func (habit Habit) String() string {
	out, err := json.MarshalIndent(habit, "", "   ")
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}
