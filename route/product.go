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

	// swagger:route GET /product/products Product GetAllProductsQuery
	//
	// Get all products that have been requested by user
	//
	// Responses:
	//  200: GetAllProductsResponse
	//  400: swaggerResponse
	routes.GET("/products", httpHandler.GetAllProducts)

	// swagger:route PATCH /product/update-status Product UpdateProductsPayload
	//
	// Update product status with either rejected or accepted
	//
	// Responses:
	//  200: swaggerResponse
	//  400: swaggerResponse
	routes.PATCH("/update-status", httpHandler.UpdateProductsStatus)

	// swagger:route DELETE /order/product-delete Product DeleteProductsPayload
	//
	// Delete products from database
	//
	// Responses:
	//  200: swaggerResponse
	//  400: swaggerResponse
	routes.DELETE("/product-delete", httpHandler.DeleteProducts)
}
