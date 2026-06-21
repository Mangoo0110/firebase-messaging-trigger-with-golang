package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"order-push-demo/logger"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

type RegisterDeviceRequest struct {
	UserID string `form:"user_id"`
	Token  string `form:"token"`
}


var fcmClient *messaging.Client


var deviceTokens = map[string]string{}


func InitFirebase() {
	ctx := context.Background()

	opt := option.WithCredentialsFile("serviceAccountKey.json")

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing firebase app %v", err)
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error initializing firebase messaging client: %v", err)
	}

	fcmClient = client

	log.Println("Firebase initialized successfully")
}


func SendPush(merchantId string, title string, body string) error {
	ctx := context.Background()

	token, found := deviceTokens[merchantId]

	logger.AppLogger.Debug.Println("merchant("+merchantId+") token: "+token+"", )

	if !found {
		logger.AppLogger.Debug.Println("Token not found!")
		return fmt.Errorf("token not found for merchant id: %s", merchantId)
	}

	message := &messaging.Message{
		Token: token,
		Notification: &messaging.Notification{
			Title: title,
			Body: body,
			// ImageURL: ,
		},

		Data: map[string]string {
			"type": "new_order",
		},
	}

	response, err :=  fcmClient.Send(ctx, message)
	if err != nil {
		return  err
	}

	logger.AppLogger.Debug.Println("Successfully sent message:", response)
	return nil
}



func RegisterDevice(c *gin.Context) {

	var req RegisterDeviceRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	deviceTokens[req.UserID] = req.Token

	logger.AppLogger.Debug.Println("Merchant id: "+req.UserID+",   Token: "+req.Token+"")

	c.JSON(http.StatusOK, gin.H{
		"message": "Device registered!",
	})
}



