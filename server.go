package main

import (
	"github.com/briannbig/habitify/handlers"
	"github.com/gin-gonic/gin"
)

var (
	habitHandler = handlers.NewHabitHandler()
	userHandler  = handlers.NewUserHandler()
)

func main() {
	r := gin.Default()

	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/", userHandler.Create)
		userRoutes.GET("/", userHandler.GetAll)
		userRoutes.PUT("/:id", userHandler.Update)
		userRoutes.DELETE("/:id", userHandler.Delete)
	}

	habitRoutes := r.Group("/habits")
	{
		habitRoutes.POST("/", habitHandler.Create)
		habitRoutes.GET("/", habitHandler.GetAll)
		habitRoutes.PUT("/:id", habitHandler.Update)
		habitRoutes.DELETE("/:id", habitHandler.Delete)
	}

	r.Run(":5050")
}
