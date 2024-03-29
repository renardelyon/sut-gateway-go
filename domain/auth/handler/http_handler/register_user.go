package http_handler

import (
	"context"
	"net/http"
	"sut-gateway-go/domain/auth/payload"
	"sut-gateway-go/helpers/http_response"
	authpb "sut-gateway-go/pb/auth"

	"sut-gateway-go/helpers/timestamp"

	"github.com/gin-gonic/gin"
)

func (h *handler) RegisterUser(ctx *gin.Context) {
	body := payload.RegisterUserPayload{}

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
	res, _ := h.auth.RegisterUser(context.Background(), &authpb.UserRegisterRequest{
		AdminId:  body.AdminId,
		Username: body.Username,
		Name:     body.Name,
		Password: body.Password,
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
		Message:    "Data successfuly registered",
		Status:     "OK",
		Timestamp:  timestamp.GetNow(),
	}

	ctx.JSON(http.StatusOK, response)
}
