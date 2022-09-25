package controllers

import (
	"github.com/InformaticsResearchCenter/chatservice/models"
	"github.com/InformaticsResearchCenter/chatservice/modules"
	"github.com/gin-gonic/gin"
)

func GetReply(ctx *gin.Context) {
	var chat models.Chat
	ctx.BindJSON(&chat)
	chat.Status = "true"
	chat.Message = modules.Chat()
	ctx.JSON(200, &chat)
}
