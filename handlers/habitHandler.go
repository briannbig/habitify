package handlers

import (
	"net/http"
	"strconv"

	"github.com/briannbig/habitify/model"
	"github.com/gin-gonic/gin"
)

type habitHandler struct {
}

func NewHabitHandler() appHandler {
	return habitHandler{}
}

func (h habitHandler) Create(ctx *gin.Context) {
	var habit model.Habit
	if err := ctx.ShouldBindJSON(&habit); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdHabit, err := habitRepo.Create(habit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdHabit)
}

func (h habitHandler) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid habit ID"})
		return
	}

	if err := habitRepo.Delete(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h habitHandler) GetAll(ctx *gin.Context) {
	habits := habitRepo.GetAll()
	ctx.JSON(http.StatusOK, habits)
}

func (h habitHandler) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid habit ID"})
		return
	}

	var habit model.Habit
	if err := ctx.ShouldBindJSON(&habit); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedHabit, err := habitRepo.Update(uint(id), habit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedHabit)
}
