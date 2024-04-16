package main

import (
	"github.com/briannbig/habitify/internal/presentation/handlers"
	"github.com/briannbig/habitify/internal/presentation/middlewares"
	"github.com/gin-gonic/gin"
)

var (
	habitHandler = handlers.NewHabitHandler()
	userHandler  = handlers.NewUserHandler()
)

func main() {
	r := gin.Default()

	userRoutes := r.Group("/users").Use(middlewares.AuthMiddleware())
	{
		userRoutes.POST("/", userHandler.Create)
		userRoutes.GET("/", userHandler.GetAll)
		userRoutes.PUT("/:id", userHandler.Update)
		userRoutes.DELETE("/:id", userHandler.Delete)
	}

	habitRoutes := r.Group("/habits").Use(middlewares.AuthMiddleware())
	{
		habitRoutes.POST("/", habitHandler.Create)
		habitRoutes.GET("/", habitHandler.GetAll)
		habitRoutes.PUT("/:id", habitHandler.Update)
		habitRoutes.DELETE("/:id", habitHandler.Delete)
	}

	r.POST("/register", userHandler.Create)
	r.POST("/login", handlers.LoginHandler)

	r.Run(":5050")
}
