package controller

import (
	"strconv"

	"github.com/KaiYanagisawa/fib_api/model"
	"github.com/gin-gonic/gin"
)

func ShowFibNumber(c *gin.Context) {
	num, _ := strconv.Atoi(c.Query("n"))
	if num <= 0 {
		BadRequest(c)
		return
	}

	fib_num := model.CalcFib(num)

	c.JSON(200, gin.H{
		"result": fib_num,
	})
}