package http_handler

import (
	"context"
	"net/http"
	"strings"
	"sut-gateway-go/domain/product/response"
	"sut-gateway-go/helpers/http_response"
	"sut-gateway-go/helpers/timestamp"
	"sut-gateway-go/pb/auth"
	productpb "sut-gateway-go/pb/product"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetAllProducts(ctx *gin.Context) {
	select {
	case <-ctx.Done():
		return
	case <-ctx.Request.Cancel:
		return
	default:
	}

	value, _ := ctx.Get("userInfo")
	userInfo := value.(*auth.UserInfo)

	query, ok := ctx.GetQuery("status")

	if !ok {
		response := http_response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Query string needed doesn't exist",
			Status:     "Bad request",
			Timestamp:  timestamp.GetNow(),
		}
		ctx.JSON(http.StatusUnauthorized, response)
		return
	}

	res, _ := h.productClient.GetAllProducts(context.Background(), &productpb.GetAllProductsRequest{
		Role:   productpb.Role(userInfo.Role),
		Id:     userInfo.Id,
		Status: productpb.Status(productpb.Status_value[strings.ToUpper(query)]),
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

	resp, err := response.NewGetAllProducts(res)
	if err != nil {
		response := http_response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Status:     "Bad Gateway",
			Timestamp:  timestamp.GetNow(),
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := http_response.Response{
		StatusCode: http.StatusOK,
		Message:    "Products data",
		Status:     "OK",
		Timestamp:  timestamp.GetNow(),
		Data:       resp.ProductInfos,
	}

	ctx.JSON(http.StatusOK, response)
}
