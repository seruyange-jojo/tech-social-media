package main

import (

	"tech-social-media/database"
	"tech-social-media/routes"

	"github.com/gin-gonic/gin"
)

func main(){
	database.ConnectDB()
	
	server := gin.Default()

	routes.SetupRouter(server)

	server.Run(":8080")

}