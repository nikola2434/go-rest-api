package main

import (
	"log"
	Config "main/config"

	"net/http"

	"github.com/gin-gonic/gin"
)

var config Config.Config

func init() {
	var err error
	config, err = Config.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

    Config.ConnectDB(&config)
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
