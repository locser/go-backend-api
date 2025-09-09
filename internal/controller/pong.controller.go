package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	name := c.DefaultQuery("name", "guest")
	c.JSON(http.StatusOK, gin.H{
		"message": name,
	})
}
