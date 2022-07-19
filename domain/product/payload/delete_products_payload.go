package payload

import (
	productpb "sut-gateway-go/pb/product"

	"github.com/gin-gonic/gin"
)

type DeleteProductsPayload struct {
	UserAndProductsIds []UserAndProduct `json:"userAndProductsIds"`
}
type DeleteProductsGrpcPayload struct {
	UserAndProductsIds []*productpb.UserAndProduct
}

type UserAndProduct struct {
	UserId    string `json:"userId"`
	ProductId string `json:"productId"`
}

func NewDeleteProductsPayload(ctx *gin.Context) (DeleteProductsGrpcPayload, error) {
	var req DeleteProductsPayload
	var body DeleteProductsGrpcPayload

	err := ctx.BindJSON(&req)
	if err != nil {
		return body, err
	}

	var userAndProducts = make([]*productpb.UserAndProduct, len(req.UserAndProductsIds))

	for i, userProduct := range req.UserAndProductsIds {
		userAndProducts[i] = &productpb.UserAndProduct{
			UserId:    userProduct.UserId,
			ProductId: userProduct.ProductId,
		}
	}

	return DeleteProductsGrpcPayload{
		UserAndProductsIds: userAndProducts,
	}, nil
}
