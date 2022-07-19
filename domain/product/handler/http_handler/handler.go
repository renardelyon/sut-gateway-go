package http_handler

import (
	productpb "sut-gateway-go/pb/product"

	"github.com/gin-gonic/gin"
)

type handler struct {
	productClient productpb.ProductServiceClient
}

type Handler interface {
	DeleteProducts(ctx *gin.Context)
	GetAllProducts(ctx *gin.Context)
	UpdateProductsStatus(ctx *gin.Context)
}

func NewProductHandler(productClient productpb.ProductServiceClient) Handler {
	return &handler{
		productClient: productClient,
	}
}
