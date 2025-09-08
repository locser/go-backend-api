package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	v2 := r.Group("/v2")
	{
		v2.GET("/ping", ping)
	}

	return r
}

func ping(c *gin.Context) {
	name := c.DefaultQuery("name", "guest")
	c.JSON(http.StatusOK, gin.H{
		"message": name,
	})
}
