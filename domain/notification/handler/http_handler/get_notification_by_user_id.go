package http_handler

import (
	"context"
	"net/http"
	notifRes "sut-gateway-go/domain/notification/response"
	"sut-gateway-go/helpers/http_response"
	"sut-gateway-go/helpers/timestamp"
	"sut-gateway-go/pb/auth"
	"sut-gateway-go/pb/notification"

	"github.com/gin-gonic/gin"
)

// swagger:parameters GetNotificationByUserIdHeader
type GetNotificationByUserIdHeader struct {
	// in: header
	// required: true
	// example: JWT eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
	Authorization string
}

func (h *handler) GetNotificationByUserId(ctx *gin.Context) {
	select {
	case <-ctx.Done():
		return
	case <-ctx.Request.Cancel:
		return
	default:
	}

	userInfo, _ := ctx.Get("userInfo")

	res, _ := h.notifClient.GetNotificationByUserId(context.Background(), &notification.GetNotificationRequest{
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

	resp := notifRes.NewGetNotificationResponse(res)

	response := http_response.Response{
		StatusCode: http.StatusOK,
		Message:    "Notification Data",
		Status:     "OK",
		Timestamp:  timestamp.GetNow(),
		Data:       resp,
	}

	ctx.JSON(http.StatusOK, response)
}
