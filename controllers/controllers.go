package controllers

import (
	"github.com/InformaticsResearchCenter/chatservice/models"
	"github.com/gin-gonic/gin"
)

func GetReply(ctx *gin.Context) {
	var chat models.Chat
	ctx.BindJSON(&chat)
	chat.Status = "true"
	ctx.JSON(200, &chat)
}
