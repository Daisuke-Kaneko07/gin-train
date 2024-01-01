package main

import (
	"gin-train/controller"
)

func main() {
	router := controller.GetRouter()
	router.Run(":8080")
}