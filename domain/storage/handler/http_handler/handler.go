package http_handler

import (
	proto "sut-gateway-go/pb/storage"

	"github.com/gin-gonic/gin"
)

type handler struct {
	storage proto.StorageServiceClient
}

type Handler interface {
	AddFile(ctx *gin.Context)
}

func NewStorageHandler(storage proto.StorageServiceClient) Handler {
	return &handler{
		storage: storage,
	}
}
