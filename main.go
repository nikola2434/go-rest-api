package main

import (
	"log"
	Config "main/config"
	controller "main/controllers"
	"main/routes"

	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine
	config Config.Config

	UserController      controller.UserController
	UserRouteController routes.UserRouteController

	AuthController       controller.AuthController
	AuthRouterController routes.AuthRouter
)

func init() {
	var err error
	config, err = Config.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	Config.ConnectDB(&config)

	UserController = controller.NewUserController(Config.DB)
	UserRouteController = routes.NewUserRouteController(UserController)

	AuthController = controller.NewAuthController(Config.DB)
	AuthRouterController = routes.NewAuthRouter(AuthController)
	server = gin.Default()
}

func main() {

	router := server.Group("/api")
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Welcome!",
		})
	})

	UserRouteController.UserRoute(router)
	AuthRouterController.AuthRouter(router)

	err := server.Run("localhost:" + config.ServerPort)
	if err != nil {
		panic(err)
	}

}
