package router

import (
	controller "go-backend/internal/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	userController := controller.NewUserController()

	v2 := r.Group("/v2")
	{
		v2.GET("/ping", controller.Ping)
		// controller user
		v2.GET("/user/detail", userController.GetDetailUser)
		v2.GET("/user/detail-error", userController.GetDetailUserError)
	}

	return r
}
