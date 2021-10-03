package main

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
   	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/thitiratratrat/hhor/src/router"
	docs "github.com/thitiratratrat/hhor/docs"
	"os"
)

var server *gin.Engine


func main() {
	server = gin.Default()

	docs.SwaggerInfo.Title = "HHor API"
	docs.SwaggerInfo.Description = "HHor is a mobile application that helps students to find nearby dormitories around the university and roommates by allowing the students to search for interested dorms, view dorm’s details, chat with dorm owners, and reserve the room."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router.InitRoutes(server)
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	
	port := os.Getenv("PORT")
	
	server.Run(":" + port)
}
