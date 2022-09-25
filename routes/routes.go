package routes

import (
	"github.com/InformaticsResearchCenter/chatservice/controllers"
	"github.com/gin-gonic/gin"
)

func ChatRoute(router *gin.Engine) {
	router.POST("/chat", controllers.GetReply)
	router.GET("/", controllers.HomeController)
}
