package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowFibNumber(c *gin.Context) {
	num, _ := strconv.Atoi(c.Param("num"))
	fib_num := model.calcFib(num)
}