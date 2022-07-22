package route

import (
	"log"
	"sut-gateway-go/config"
	"sut-gateway-go/domain/storage/handler/http_handler"
	storagepb "sut-gateway-go/pb/storage"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func SetupStorageRouter(r *gin.Engine, cfg *config.Config) {
	storageGrpcConn, err := grpc.Dial(
		cfg.StorageSvcUrl,
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("error when dialing grpc client")
	}

	storageGrpcService := storagepb.NewStorageServiceClient(storageGrpcConn)

	httpHandler := http_handler.NewStorageHandler(storageGrpcService)

	routes := r.Group("/storage")

	// swagger:route POST /storage/add Storage AddFilePayload
	//
	// Upload file into google drive
	//
	// Consumes:
	// - multipart/form-data
	//
	// Produces:
	// - application/json
	//
	// Responses:
	//  200: swaggerResponse
	//  400: swaggerResponse
	routes.POST("/add", httpHandler.AddFile)
}
