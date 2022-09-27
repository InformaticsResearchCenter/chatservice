package main

import (
	"github.com/InformaticsResearchCenter/chatservice/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	routes.ChatRoute(router)
	router.Run(":5000")
}
