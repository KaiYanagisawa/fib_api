package controller

import (
	"strconv"

	"github.com/KaiYanagisawa/fib_api/model"
	"github.com/gin-gonic/gin"
)

func ShowFibNumber(c *gin.Context) {
	num, _ := strconv.Atoi(c.Query("n"))
	
	// NOTE: num=250000のとき実行時間が約200ms. それ以上のリクエストは受け付けないようにする
	if num <= 0 || num > 250000{
		BadRequest(c)
		return
	}

	fib_num := model.CalcFib(num)

	c.JSON(200, gin.H{
		"result": fib_num,
	})
}