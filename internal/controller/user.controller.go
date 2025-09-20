package controller

import (
	service "go-backend/internal/service"
	"go-backend/responses"

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

	responses.SuccessResponse(c, responses.StatusOK, uc.userService.GetDetailUser())
}

func (uc *UserController) GetDetailUserError(c *gin.Context) {

	responses.ErrorResponse(c, responses.ErrCodeParamInvalid, "Test Error")
}
