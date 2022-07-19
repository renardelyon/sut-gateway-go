package route

import (
	"log"
	"sut-gateway-go/config"
	"sut-gateway-go/domain/product/handler/http_handler"
	productpb "sut-gateway-go/pb/product"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func SetupProductRouter(r *gin.Engine, cfg *config.Config) {
	productGrpcConn, err := grpc.Dial(
		cfg.ProductSvcUrl,
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("error when dialing grpc client")
	}

	productGrpcService := productpb.NewProductServiceClient(productGrpcConn)

	httpHandler := http_handler.NewProductHandler(productGrpcService)

	routes := r.Group("/product")
	routes.GET("/products", httpHandler.GetAllProducts)
	routes.PATCH("/update-status", httpHandler.UpdateProductsStatus)
	routes.DELETE("/product-delete", httpHandler.DeleteProducts)
}
