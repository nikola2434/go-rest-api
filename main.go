package main

import (
	"log"
	init_env "main/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

var config init_env.Config

func init() {
	var err error
	config, err = init_env.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}
}

func main() {

	route := gin.Default()
	route.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "qwe",
		})
	})

	err := route.Run("localhost:" + config.ServerPort)
	if err != nil {
		panic(err)
	}

}
