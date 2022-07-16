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
	routes.GET("/", httpHandler.GetNotificationByUserId)
	routes.PATCH("/", httpHandler.ResetNotificationStatusQty)
}
