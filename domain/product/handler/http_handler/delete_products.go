package http_handler

import (
	"context"
	"net/http"
	"sut-gateway-go/domain/product/payload"
	"sut-gateway-go/helpers/http_response"
	"sut-gateway-go/helpers/timestamp"
	productpb "sut-gateway-go/pb/product"

	"github.com/gin-gonic/gin"
)

func (h *handler) DeleteProducts(ctx *gin.Context) {
	select {
	case <-ctx.Done():
		return
	case <-ctx.Request.Cancel:
		return
	default:
	}

	body, err := payload.NewDeleteProductsPayload(ctx)
	if err != nil {
		response := http_response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Status:     "Bad Request",
			Timestamp:  timestamp.GetNow(),
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	res, _ := h.productClient.DeleteProducts(context.Background(), &productpb.DeleteProductsRequest{
		UserAndProductsIds: body.UserAndProductsIds,
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
		Message:    "Order product has been deleted",
		Status:     "OK",
		Timestamp:  timestamp.GetNow(),
	}

	ctx.JSON(http.StatusOK, response)
}
