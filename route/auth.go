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

	// swagger:route POST /auth/register Auth RegisterUserPayload
	//
	// Register User
	//
	// Responses:
	//  200: swaggerResponse
	routes.POST("/register", httpHandler.RegisterUser)

	// swagger:route POST /auth/register-admin Auth RegisterAdminPayload
	//
	// Register Admin
	//
	// Consumes:
	// - application/json
	//
	// Produces:
	// - application/json
	//
	// Responses:
	//  200: swaggerResponse
	//  400: swaggerResponse
	routes.POST("/register-admin", httpHandler.RegisterAdmin)

	// swagger:route POST /auth/login Auth LoginPayload
	//
	// Login User and Admin
	//
	// Consumes:
	// - application/json
	//
	// Produces:
	// - application/json
	//
	// Responses:
	//  200: LoginResponse
	//  400: swaggerResponse
	routes.POST("/login", httpHandler.Login)

	// swagger:route POST /auth/regenerate-token Auth GenerateNewTokenPayload
	//
	// Regenerate access token from refresh token
	//
	// Consumes:
	// - application/json
	//
	// Produces:
	// - application/json
	//
	// Responses:
	//  200: GenerateNewTokenResponse
	//  400: swaggerResponse
	routes.POST("/regenerate-token", httpHandler.GenerateNewToken)

	// swagger:route POST /auth/auth-token Auth NoParams
	//
	// Get user info from jwt token
	//
	// Consumes:
	// - application/json
	//
	// Produces:
	// - application/json
	//
	// Responses:
	//  200: UserInfo
	//  400: swaggerResponse
	routes.POST("/auth-token", authMidWare, func(ctx *gin.Context) {
		userInfo, _ := ctx.Get("userInfo")
		ctx.JSON(http.StatusOK, userInfo)
	})

	// swagger:route DELETE /auth/logout Auth DeleteToken
	//
	// Delete token from database
	//
	// Consumes:
	// - application/json
	//
	// Produces:
	// - application/json
	//
	// Responses:
	//  200: swaggerResponse
	//  400: swaggerResponse
	routes.DELETE("/logout", httpHandler.DeleteToken)

	routes.Use(authMidWare)

	// swagger:route GET /auth/register Auth GetUserByRoleQuery
	//
	// Get all users that has registered to the app
	//
	// Responses:
	//  200: GetUserByRoleResponse
	routes.GET("/register", httpHandler.GetUserByRole)
}
