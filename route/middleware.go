package route

import (
	"log"
	"sut-gateway-go/config"
	"sut-gateway-go/domain/auth/util"
	authpb "sut-gateway-go/pb/auth"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func SetupAuthMiddleWare(r *gin.Engine, cfg *config.Config) {
	authGrpcConn, err := grpc.Dial(
		cfg.AuthSvcUrl,
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("error when dialing grpc client")
	}
	authGrpcService := authpb.NewAuthServiceClient(authGrpcConn)

	authMidWare := util.NewAuthMiddleWare(authGrpcService)

	r.Use(authMidWare)
}
