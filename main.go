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
	Server    *gin.Engine
	ConfigApp Config.Config

	UserController      controller.UserController
	UserRouteController routes.UserRouteController

	AuthController       controller.AuthController
	AuthRouterController routes.AuthRouter
)

func init() {
	var err error
	ConfigApp, err = Config.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	Config.ConnectDB(&ConfigApp)

	UserController = controller.NewUserController(Config.DB)
	UserRouteController = routes.NewUserRouteController(UserController)

	AuthController = controller.NewAuthController(Config.DB)
	AuthRouterController = routes.NewAuthRouter(AuthController)
	Server = gin.Default()
}

func main() {

	router := Server.Group("/api")
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Welcome!",
		})
	})

	UserRouteController.UserRoute(router)
	AuthRouterController.AuthRouter(router)

	err := Server.Run("localhost:" + ConfigApp.ServerPort)
	if err != nil {
		panic(err)
	}

}
