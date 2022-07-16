package route

import (
	"log"
	"net/http"
	"sut-gateway-go/config"

	"sut-gateway-go/domain/auth/handler/http_handler"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"sut-gateway-go/domain/auth/util"
	authpb "sut-gateway-go/pb/auth"
)

func SetupAuthRouter(r *gin.Engine, cfg *config.Config) {
	authGrpcConn, err := grpc.Dial(
		cfg.AuthSvcUrl,
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("error when dialing grpc client")
	}

	authGrpcService := authpb.NewAuthServiceClient(authGrpcConn)

	authMidWare := util.NewAuthMiddleWare(authGrpcService)

	httpHandler := http_handler.NewAuthHandler(authGrpcService)

	routes := r.Group("/auth")
	routes.POST("/register", httpHandler.RegisterUser)
	routes.POST("/register-admin", httpHandler.RegisterAdmin)
	routes.POST("/login", httpHandler.Login)
	routes.POST("/regenerate-token", httpHandler.GenerateNewToken)
	routes.POST("/auth-token", authMidWare, func(ctx *gin.Context) {
		userInfo, _ := ctx.Get("userInfo")
		ctx.JSON(http.StatusOK, userInfo)
	})
	routes.DELETE("/logout", httpHandler.DeleteToken)
	routes.Use(authMidWare)
	routes.GET("/register", httpHandler.GetUserByRole)
}
