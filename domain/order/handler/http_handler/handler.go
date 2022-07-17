package http_handler

import (
	orderpb "sut-gateway-go/pb/order"

	"github.com/gin-gonic/gin"
)

type handler struct {
	orderClient orderpb.OrderServiceClient
}

type Handler interface {
	GetDetailProducts(ctx *gin.Context)
	GetProductsToOrderByKeyword(ctx *gin.Context)
	CreateOrder(ctx *gin.Context)
	CreateOrderBulk(ctx *gin.Context)
}

func NewNotificationHandler(orderClient orderpb.OrderServiceClient) Handler {
	return &handler{
		orderClient: orderClient,
	}
}
