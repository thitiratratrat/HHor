package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thitiratratrat/hhor/src/router"
)

var server *gin.Engine

func init() {
	server = gin.Default()

	router.InitRoutes(server)
}

func main() {
	server.Run("localhost:8081")
}
