package main

import (
	"github.com/rishikeshbedre/nats-api-server/apis"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/user", apis.ShowUsers)
	router.POST("/user", apis.AddUser)
	router.DELETE("/user", apis.DeleteUser)

	router.POST("/topic", apis.AddTopic)
	router.DELETE("/topic", apis.DeleteTopic)

	router.POST("/reload", apis.DownloadConfiguration)

	router.Run(":6060") // listen and serve on 6060
}
