package http_handler

import (
	"context"
	"net/http"
	"sut-gateway-go/domain/auth/payload"
	"sut-gateway-go/helpers/http_response"
	"sut-gateway-go/helpers/timestamp"
	"sut-gateway-go/pb/auth"

	"github.com/gin-gonic/gin"
)

func (h *handler) Login(ctx *gin.Context) {
	body := payload.LoginPayload{}

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

	res, _ := h.auth.Login(context.Background(), &auth.LoginRequest{
		Username: body.Username,
		Password: body.Password,
	})

	if res.Error != "" {
		response := http_response.Response{
			StatusCode: http.StatusBadGateway,
			Message:    res.Error,
			Status:     "Bad Request",
			Timestamp:  timestamp.GetNow(),
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := http_response.Response{
		StatusCode: http.StatusOK,
		Message:    "OK",
		Status:     "OK",
		Timestamp:  timestamp.GetNow(),
		Data:       res,
	}

	ctx.JSON(http.StatusOK, response)
}
