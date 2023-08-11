package main

import (
	"github.com/KaiYanagisawa/fib_api/controller"
)

func main() {
	router := controller.GetRouter()
	router.Run(":8080")
}