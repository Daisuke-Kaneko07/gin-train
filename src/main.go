package main

import (
	"controller"
)

func main() {
	router := controller.GetRouter()
	router.Run(":8080")
}