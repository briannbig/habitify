package handlers

import (

	"github.com/briannbig/habitify/repository"
	"github.com/gin-gonic/gin"
)

var (
	habitRepo = repository.NewHabitRepository()
	userRepo  = repository.NewUserRepository()
)

type appHandler interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
