package http_handler

import (
	"context"
	"net/http"
	"sut-gateway-go/helpers/http_response"
	"sut-gateway-go/helpers/timestamp"
	"sut-gateway-go/pb/auth"
	orderpb "sut-gateway-go/pb/order"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateOrderBulk(ctx *gin.Context) {
	select {
	case <-ctx.Done():
		return
	case <-ctx.Request.Cancel:
		return
	default:
	}

	ui, _ := ctx.Get("userInfo")
	userInfo := ui.(*auth.UserInfo)

	res, _ := h.orderClient.CreateOrderBulk(context.Background(), &orderpb.CreateOrderBulkRequest{
		AdminId:  userInfo.AdminId,
		Fullname: userInfo.Name,
		Username: userInfo.Username,
		UserId:   userInfo.Id,
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
		Message:    "Bulk order has been created",
		Status:     "OK",
		Timestamp:  timestamp.GetNow(),
	}

	ctx.JSON(http.StatusOK, response)
}
