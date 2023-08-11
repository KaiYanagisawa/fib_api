package controller

import (
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/fib", ShowFibNumber)

	r.NoRoute(NotFound)

	r.HandleMethodNotAllowed = true
    r.NoMethod(NotImplemented)
	
	return r
}
