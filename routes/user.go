package routes

import (
	"github.com/InformaticsResearchCenter/chatservice/controller"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/", controller.UserController)
}
