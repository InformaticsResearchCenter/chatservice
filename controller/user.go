package controller

import (
	"github.com/InformaticsResearchCenter/chatservice/models"
	"github.com/gin-gonic/gin"
)

func UserController(ctx *gin.Context) {
	users := []models.User{}
	confi.DB.Finf(&users)
	ctx.JSON(200, &users)
}
