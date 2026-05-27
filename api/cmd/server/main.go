package main

import (
	
	"log"

	"chat-platform-api/configs"
	"chat-platform-api/pkg/database"

	"github.com/gin-gonic/gin"
	"chat-platform-api/routes"
)
func main(){
	configs.LoadEnv()
	err:= database.ConnectMongo(
		configs.GetEnv("MONGO_URI"),
		configs.GetEnv("DATABASE_NAME"),
	)

	if err != nil{
		log.Fatal(err)
	}
	r:= gin.Default()

	routes.SetupRoutes(r)
	r.Run(":"+configs.GetEnv("PORT"))
}