package controller

import (
	"strconv"

	"github.com/KaiYanagisawa/fib_api/model"
	"github.com/gin-gonic/gin"
)

func ShowFibNumber(c *gin.Context) {
	num, _ := strconv.Atoi(c.Query("n"))
	fib_num := model.CalcFib(num)
	c.JSON(200, gin.H{
		"result": fib_num,
	})
}