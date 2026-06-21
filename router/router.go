package router

import (
	"order-push-demo/controller"
	"github.com/gin-gonic/gin"
)

func SetRouter(route *gin.Engine){
	route.POST("devices/register", controller.RegisterDevice)
	route.POST("order/place",controller.PlaceOrder)
}