package http_handler

import (
	"sut-gateway-go/pb/auth"

	"github.com/gin-gonic/gin"
)

type handler struct {
	auth auth.AuthServiceClient
}

type Handler interface {
	RegisterUser(ctx *gin.Context)
	RegisterAdmin(ctx *gin.Context)
	DeleteToken(ctx *gin.Context)
	GenerateNewToken(ctx *gin.Context)
	Login(ctx *gin.Context)
	GetUserByRole(ctx *gin.Context)
}

func NewAuthHandler(auth auth.AuthServiceClient) Handler {
	return &handler{
		auth: auth,
	}
}
