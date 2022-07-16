package http_handler

import (
	notifpb "sut-gateway-go/pb/notification"

	"github.com/gin-gonic/gin"
)

type handler struct {
	notifClient notifpb.NotificationServiceClient
}

type Handler interface {
	GetNotificationByUserId(ctx *gin.Context)
	ResetNotificationStatusQty(ctx *gin.Context)
}

func NewNotificationHandler(notifClient notifpb.NotificationServiceClient) Handler {
	return &handler{
		notifClient: notifClient,
	}
}
