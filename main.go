package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	route := gin.Default()
	route.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "оикы123123",
		})
	})

	err := route.Run("localhost:3001")
	if err != nil {
		panic(err)
	}

}
