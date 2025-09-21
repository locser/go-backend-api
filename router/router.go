package router

import (
	"fmt"
	controller "go-backend/internal/controller"
	"go-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

// gin vs middleware: docs/images/go_middleware.png

func AA() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		fmt.Println("Before AA:")
		ctx.Next()
		fmt.Println("After AA:")

	}

}

func BB() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		fmt.Println("Before BB:")
		ctx.Next()
		fmt.Println("After BB:")

	}

}

func CC(ctx *gin.Context) {
	fmt.Println("Before CC:")
	ctx.Next()
	fmt.Println("After CC:")
}

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.AuthenMiddleware(), AA(), BB(), CC)

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
