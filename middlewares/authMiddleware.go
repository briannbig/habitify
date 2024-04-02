package middlewares

import (
	"github.com/briannbig/habitify/model"
	"github.com/briannbig/habitify/repository"
	"github.com/gin-gonic/gin"
)

var userRepository repository.Repository[model.User] = repository.NewUserRepository()

func AuthMiddleware() gin.HandlerFunc {

	users := userRepository.GetAll()

	accounts := map[string]string{}

	for _, user := range users {
		accounts[user.Email] = user.Password
	}

	return gin.BasicAuth(accounts)
}
