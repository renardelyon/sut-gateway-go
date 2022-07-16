package util

import (
	"context"
	"net/http"
	"strings"
	"sut-gateway-go/helpers/http_response"
	"sut-gateway-go/helpers/timestamp"
	"sut-gateway-go/pb/auth"

	"github.com/gin-gonic/gin"
)

func NewAuthMiddleWare(authClient auth.AuthServiceClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.Request.Header.Get("authorization")

		if authorization == "" {
			response := http_response.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    "JWT Token is not valid",
				Status:     "Unauthorized",
				Timestamp:  timestamp.GetNow(),
			}
			ctx.JSON(http.StatusUnauthorized, response)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token := strings.Split(authorization, " ")[1]

		if token == "" {
			response := http_response.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    "JWT Token is not valid",
				Status:     "Unauthorized",
				Timestamp:  timestamp.GetNow(),
			}
			ctx.JSON(http.StatusUnauthorized, response)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		res, _ := authClient.Authenticate(context.Background(), &auth.AuthenticateRequest{
			Token: token,
		})

		if res.Error != "" || res.Status != http.StatusOK {
			response := http_response.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    "JWT Token is not valid",
				Status:     "Unauthorized",
				Timestamp:  timestamp.GetNow(),
			}
			ctx.JSON(http.StatusUnauthorized, response)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set("userInfo", res.Userinfo)
		ctx.Next()
	}
}
