package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", helloWorld)
	r.Run(":80")
	// 發布測試
	//
}

func helloWorld(c *gin.Context) {
	fmt.Println("Hello World!")
}

func unuseFunc() {
	fmt.Println("This is an unuse function.")
}
