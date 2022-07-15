package http_handler

import (
	"context"
	"net/http"
	"sut-gateway-go/domain/auth/payload"
	"sut-gateway-go/helpers/http_response"
	"sut-gateway-go/helpers/timestamp"
	"sut-gateway-go/pb/auth"

	"sut-gateway-go/domain/auth/response"

	"github.com/gin-gonic/gin"
)

func (h *handler) GenerateNewToken(ctx *gin.Context) {
	body := payload.GenerateNewTokenPayload{}

	if err := ctx.BindJSON(&body); err != nil {
		response := http_response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Status:     "Bad Request",
			Timestamp:  timestamp.GetNow(),
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	res, _ := h.auth.GenerateNewToken(context.Background(), &auth.GenerateNewTokenRequest{
		RefreshToken: body.Token,
	})

	if res.Error != "" {
		response := http_response.Response{
			StatusCode: http.StatusBadGateway,
			Message:    res.Error,
			Status:     "Bad Gateway",
			Timestamp:  timestamp.GetNow(),
		}
		ctx.JSON(http.StatusBadGateway, response)
		return
	}

	response := http_response.Response{
		StatusCode: http.StatusOK,
		Message:    "Status OK",
		Status:     "OK",
		Timestamp:  timestamp.GetNow(),
		Data:       response.NewGenerateTokenResponse(res.NewToken),
	}

	ctx.JSON(http.StatusOK, response)
}
