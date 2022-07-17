package http_handler

import (
	"context"
	"net/http"
	orderRes "sut-gateway-go/domain/order/response"
	"sut-gateway-go/helpers/http_response"
	"sut-gateway-go/helpers/timestamp"

	orderpb "sut-gateway-go/pb/order"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetProductsToOrderByKeyword(ctx *gin.Context) {
	select {
	case <-ctx.Done():
		return
	case <-ctx.Request.Cancel:
		return
	default:
	}

	query, ok := ctx.GetQuery("keyword")

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

	res, _ := h.orderClient.GetProductsToOrderByKeyword(context.Background(), &orderpb.GetProductsToOrderRequest{
		Keyword: query,
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

	resp, err := orderRes.NewGetProductsToOrderResponse(res)

	if err != nil {
		response := http_response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Status:     "Internal Server Error",
			Timestamp:  timestamp.GetNow(),
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := http_response.Response{
		StatusCode: http.StatusOK,
		Message:    "Data fetch from bukalapak",
		Status:     "OK",
		Timestamp:  timestamp.GetNow(),
		Data:       resp,
	}

	ctx.JSON(http.StatusOK, response)
}
