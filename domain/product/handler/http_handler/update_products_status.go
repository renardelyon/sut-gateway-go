package http_handler

import (
	"context"
	"log"
	"net/http"
	"sut-gateway-go/domain/product/payload"
	"sut-gateway-go/helpers/http_response"
	"sut-gateway-go/helpers/timestamp"
	"sut-gateway-go/pb/auth"
	"sut-gateway-go/pb/product"

	"github.com/gin-gonic/gin"
)

func (h *handler) UpdateProductsStatus(ctx *gin.Context) {
	select {
	case <-ctx.Done():
		return
	case <-ctx.Request.Cancel:
		return
	default:
	}

	value, _ := ctx.Get("userInfo")
	userInfo := value.(*auth.UserInfo)

	body, err := payload.NewUpdateProductsPayload(ctx)
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

	res, _ := h.productClient.UpdateProductStatus(context.Background(), &product.UpdateProductStatusRequest{
		Role:         product.Role(userInfo.Role),
		Status:       body.Status,
		UserProducts: body.UserAndProductIds,
	})

	log.Println("res", res)

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
		Message:    "Products status successfully updated",
		Status:     "OK",
		Timestamp:  timestamp.GetNow(),
	}

	ctx.JSON(http.StatusOK, response)
}
