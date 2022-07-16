package http_handler

import (
	"context"
	"net/http"
	"sut-gateway-go/helpers/http_response"
	"sut-gateway-go/helpers/timestamp"
	"sut-gateway-go/pb/auth"
	notifpb "sut-gateway-go/pb/notification"

	"github.com/gin-gonic/gin"
)

func (h *handler) ResetNotificationStatusQty(ctx *gin.Context) {
	select {
	case <-ctx.Done():
		return
	case <-ctx.Request.Cancel:
		return
	default:
	}

	userInfo, _ := ctx.Get("userInfo")

	res, _ := h.notifClient.ResetNotificationStatusQty(context.Background(), &notifpb.ResetNotificationRequest{
		UserId: userInfo.(*auth.UserInfo).Id,
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
		Message:    "Notification Data",
		Status:     "OK",
		Timestamp:  timestamp.GetNow(),
	}

	ctx.JSON(http.StatusOK, response)
}
