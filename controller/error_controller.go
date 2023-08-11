package controller

import (
	"github.com/gin-gonic/gin"
)

func NotFound(c *gin.Context) {
	c.JSON(404, gin.H{
		"status":  404,
		"message": "Not Found",
	})
}

func NotImplemented(c *gin.Context) {
	c.JSON(501, gin.H{
		"status":  501,
		"message": "Not Implemented",
	})
}
