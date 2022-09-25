package controllers

import (
	"github.com/InformaticsResearchCenter/chatservice/models"
	"github.com/gin-gonic/gin"
)

func UserController(ctx *gin.Context) {
	rolly := models.User{Id: 113040087, Nama: "Ganteng rolly"}
	kaka := models.User{Id: 01212323, Nama: "kaka"}
	users := []models.User{}
	users = append(users, rolly, kaka)
	ctx.JSON(200, &users)
}
