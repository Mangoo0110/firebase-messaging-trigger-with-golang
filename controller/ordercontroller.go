package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PlaceOrderRequest struct {
	OrderID    string `form:"order_id"`
	CustomerID string `form:"customer_id"`
	MerchantID string `form:"merchant_id"`
}

func PlaceOrder(c *gin.Context) {
	var req PlaceOrderRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := SendPush(req.MerchantID, "New Order Received", "Order "+req.OrderID+" has been placed.")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "failed to send push notification",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "order placed and notification sent",
		"order_id":    req.OrderID,
		"merchant_id": req.MerchantID,
	})

}