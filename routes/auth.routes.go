package routes

import (
	controller "main/controllers"

	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
	AuthRouterController controller.AuthController
}

func NewAuthRouter(AuthRouterController controller.AuthController) AuthRouter {
	return AuthRouter{AuthRouterController}
}

func (ac *AuthRouter) AuthRouter(ctx *gin.RouterGroup) {
	router := ctx.Group("auth")
	router.GET("qwe")
}
