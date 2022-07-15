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

func (h *handler) GetUserByRole(ctx *gin.Context) {
	userInfo, _ := ctx.Get("userInfo")

	if userInfo.(*auth.UserInfo).Role != auth.Role_ADMIN {
		response := http_response.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "Unauthorized",
			Status:     "Unauthorized",
			Timestamp:  timestamp.GetNow(),
		}
		ctx.JSON(http.StatusUnauthorized, response)
		return
	}

	body := payload.GetUserPayload{}

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

	res, _ := h.auth.GetUserByRole(context.Background(), &auth.GetUserByRoleRequest{
		Role: body.Role,
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
		Message:    "OK",
		Status:     "OK",
		Timestamp:  timestamp.GetNow(),
		Data:       res.UserAdminInfo,
	}

	ctx.JSON(http.StatusOK, response)
}
