package routes

import (
	controller "main/controllers"

	"github.com/gin-gonic/gin"
)

type UserRouteController struct {
	userController controller.UserController
}

func NewUserRouteController(userController controller.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (uc *UserRouteController) UserRoute(rg *gin.RouterGroup) {
	router := rg.Group("users")
	router.GET("/profile", uc.userController.GetProfile)
}
