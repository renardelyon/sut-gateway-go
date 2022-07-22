package route

import (
	"log"
	"sut-gateway-go/config"

	"sut-gateway-go/domain/notification/handler/http_handler"
	notifpb "sut-gateway-go/pb/notification"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func SetupNotificationRouter(r *gin.Engine, cfg *config.Config) {
	notifGrpcConn, err := grpc.Dial(
		cfg.NotifSvcUrl,
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("error when dialing grpc client")
	}

	notifGrpcService := notifpb.NewNotificationServiceClient(notifGrpcConn)

	httpHandler := http_handler.NewNotificationHandler(notifGrpcService)

	routes := r.Group("/notification")

	// swagger:route GET /notification Notification GetNotificationByUserIdHeader
	//
	// Get notification data by user id
	//
	// Responses:
	//  200: GetNotificationResponse
	//  400: swaggerResponse
	routes.GET("/", httpHandler.GetNotificationByUserId)

	// swagger:route PATCH /notification Notification ResetNotificationStatusQtyHeader
	//
	// Reset notification status quantity
	//
	// Responses:
	//  200: swaggerResponse
	//  400: swaggerResponse
	routes.PATCH("/", httpHandler.ResetNotificationStatusQty)
}
