package payload

import (
	"strings"
	productpb "sut-gateway-go/pb/product"

	"github.com/gin-gonic/gin"
)

type UserAndProductId struct {
	UserId    string `json:"userId"`
	ProductId string `json:"productId"`
}

type UpdateProductsPayload struct {
	Status            string             `json:"status"`
	UserAndProductIds []UserAndProductId `json:"userAndProductIds"`
}

type UpdateProductsGrpcPayload struct {
	Status            productpb.Status
	UserAndProductIds []*productpb.UserAndProduct
}

func NewUpdateProductsPayload(ctx *gin.Context) (UpdateProductsGrpcPayload, error) {
	var req UpdateProductsPayload
	var body UpdateProductsGrpcPayload

	err := ctx.BindJSON(&req)
	if err != nil {
		return body, err
	}

	var userAndProducts = make([]*productpb.UserAndProduct, len(req.UserAndProductIds))

	for i, userProduct := range req.UserAndProductIds {
		userAndProducts[i] = &productpb.UserAndProduct{
			UserId:    userProduct.UserId,
			ProductId: userProduct.ProductId,
		}
	}

	return UpdateProductsGrpcPayload{
		Status:            productpb.Status(productpb.Status_value[strings.ToUpper(req.Status)]),
		UserAndProductIds: userAndProducts,
	}, nil
}
