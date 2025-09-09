package controller

import (
	service "go-backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetDetailUser(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": uc.userService.GetDetailUser(),
	})
}
