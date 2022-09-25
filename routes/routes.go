package routes

import (
	"github.com/InformaticsResearchCenter/chatservice/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/", controllers.UserController)
}
