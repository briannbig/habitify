package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/briannbig/habitify/repository"
	"github.com/briannbig/habitify/util"
	"github.com/gin-gonic/gin"
)

func LoginHandler(ctx *gin.Context) {
	var request requestLogin
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("request: %s\n", request)

	userRepo := repository.UserRepository{}

	user, _ := userRepo.FindUserByEmail(request.Username)

	if user.Password != request.Password {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "wrong password"})
		return
	}

	token, err := util.GenerateUserToken(user.Email, user.Role)
	if err != nil {
		errStr := fmt.Sprintf("could not generate jwt token for user --- %s", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errStr})
		return
	}

	ctx.JSON(http.StatusOK, responseLogin{Prefix: "Bearer", Token: token})

}

type (
	requestLogin struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	responseLogin struct {
		Prefix string
		Token  string
	}
)
