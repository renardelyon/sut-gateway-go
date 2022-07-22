package route

import (
	"log"
	"sut-gateway-go/config"
	"sut-gateway-go/domain/order/handler/http_handler"
	orderpb "sut-gateway-go/pb/order"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func SetupOrderRouter(r *gin.Engine, cfg *config.Config) {
	orderGrpcConn, err := grpc.Dial(
		cfg.OrderSvcUrl,
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("error when dialing grpc client")
	}

	orderGrpcService := orderpb.NewOrderServiceClient(orderGrpcConn)

	httpHandler := http_handler.NewNotificationHandler(orderGrpcService)

	routes := r.Group("/order")

	// swagger:route GET /order/products-by-keyword Order GetProductsToOrderByKeywordQuery
	//
	// Get products from bukalapak
	//
	// Responses:
	//  200: GetProductsToOrderResponse
	//  400: swaggerResponse
	routes.GET("/products-by-keyword", httpHandler.GetProductsToOrderByKeyword)

	// swagger:route GET /order/detail-product Order GetDetailProductsQuery
	//
	// Get product detail
	//
	// Responses:
	//  200: GetDetailProductsResponse
	//  400: swaggerResponse
	routes.GET("/detail-product", httpHandler.GetDetailProducts)

	// swagger:route POST /order Order CreateOrderPayload
	//
	// Create single order
	//
	// Responses:
	//  200: swaggerResponse
	//  400: swaggerResponse
	routes.POST("/", httpHandler.CreateOrder)

	// swagger:route POST /order/bulk Order CreateOrderBulkHeader
	//
	// Create bulk order
	//
	// Responses:
	//  200: swaggerResponse
	//  400: swaggerResponse
	routes.POST("/bulk", httpHandler.CreateOrderBulk)
}
