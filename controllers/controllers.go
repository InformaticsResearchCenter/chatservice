package controllers

import (
	"github.com/InformaticsResearchCenter/chatservice/models"
	"github.com/gin-gonic/gin"
)

func ChatController(ctx *gin.Context) {
	chat := models.Chat{Status: "true", Message: "Ganteng rolly"}
	ctx.JSON(200, &chat)
}
