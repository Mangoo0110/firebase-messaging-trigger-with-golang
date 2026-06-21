package main

import (
	"order-push-demo/config"
	"order-push-demo/controller"
	"order-push-demo/logger"
	"order-push-demo/router"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func init(){
	err := godotenv.Load()
	if (err != nil) {
		logger.AppLogger.Error.Println("Failed to load env file : ",err)
		os.Exit(1)
	}
	config.LoadConfig()
}


func main() {
	controller.InitFirebase()
	
	routers := gin.Default()

	router.SetRouter(routers)

	routers.Run(":8080")

	logger.AppLogger.Info.Println(config.ORDER_CODE)
}