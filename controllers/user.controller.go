package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(DB *gorm.DB) UserController {
	return UserController{DB}
}

func (uc *UserController) GetProfile(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": "profile"})
}
