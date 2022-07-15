package http_handler

import (
	"context"
	"net/http"
	"sut-gateway-go/domain/auth/payload"
	"sut-gateway-go/helpers/http_response"
	"sut-gateway-go/helpers/timestamp"
	authpb "sut-gateway-go/pb/auth"

	"github.com/gin-gonic/gin"
)

func (h *handler) RegisterAdmin(ctx *gin.Context) {
	body := payload.RegisterAdminPayload{}

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

	res, _ := h.auth.RegisterAdmin(context.Background(), &authpb.AdminRegisterRequest{
		Name:     body.Name,
		Username: body.Username,
		Password: body.Password,
		AdminKey: body.AdminKey,
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
