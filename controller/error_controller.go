package controller

import (
	"github.com/gin-gonic/gin"
)

func NoRouter(c *gin.Context) {
	c.JSON(404, gin.H{
		"status": 404,
		"message": "Not Found",
	})
}