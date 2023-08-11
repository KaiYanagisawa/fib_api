package main

import (
	"fib_api/controller"
)

func main() {
	router := controller.GetRouter()	
	router.Run(":8080")
}