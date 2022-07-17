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
	routes.GET("/products-by-keyword", httpHandler.GetProductsToOrderByKeyword)
	routes.GET("/detail-product", httpHandler.GetDetailProducts)
	routes.POST("/", httpHandler.CreateOrder)
	routes.POST("/bulk", httpHandler.CreateOrderBulk)
}
