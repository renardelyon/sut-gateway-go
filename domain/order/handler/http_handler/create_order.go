package http_handler

import (
	"context"
	"net/http"
	"sut-gateway-go/domain/order/payload"
	"sut-gateway-go/helpers/http_response"
	"sut-gateway-go/helpers/timestamp"
	"sut-gateway-go/pb/auth"

	orderpb "sut-gateway-go/pb/order"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateOrder(ctx *gin.Context) {
	select {
	case <-ctx.Done():
		return
	case <-ctx.Request.Cancel:
		return
	default:
	}

	body := payload.CreateOrderPayload{}

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
	ui, _ := ctx.Get("userInfo")
	userInfo := ui.(*auth.UserInfo)

	res, _ := h.orderClient.CreateOrder(context.Background(), &orderpb.CreateOrderRequest{
		AdminId:     userInfo.AdminId,
		Fullname:    userInfo.Name,
		Username:    userInfo.Username,
		UserId:      userInfo.Id,
		Url:         body.Url,
		ProductName: body.ProductName,
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
		Message:    "Order has been created",
		Status:     "OK",
		Timestamp:  timestamp.GetNow(),
	}

	ctx.JSON(http.StatusOK, response)
}
