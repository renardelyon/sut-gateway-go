package http_handler

import (
	"context"
	"net/http"
	"sut-gateway-go/helpers/http_response"
	"sut-gateway-go/helpers/timestamp"
	orderpb "sut-gateway-go/pb/order"

	orderRes "sut-gateway-go/domain/order/response"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetDetailProducts(ctx *gin.Context) {
	select {
	case <-ctx.Done():
		return
	case <-ctx.Request.Cancel:
		return
	default:
	}

	query, ok := ctx.GetQuery("url")

	if !ok {
		response := http_response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Query string role doesn't exist",
			Status:     "Bad request",
			Timestamp:  timestamp.GetNow(),
		}
		ctx.JSON(http.StatusUnauthorized, response)
		return
	}

	res, _ := h.orderClient.GetDetailProducts(context.Background(), &orderpb.GetDetailProductsRequest{
		Url: query,
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

	resp := orderRes.NewGetDetailProductsResponse(res)

	response := http_response.Response{
		StatusCode: http.StatusOK,
		Message:    "Data fetch from bukalapak",
		Status:     "OK",
		Timestamp:  timestamp.GetNow(),
		Data:       resp,
	}

	ctx.JSON(http.StatusOK, response)
}
