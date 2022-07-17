package response

import (
	"encoding/json"
	orderpb "sut-gateway-go/pb/order"
)

type ProductInfo struct {
	Url         string `json:"url"`
	ProductName string `json:"Name"`
}

type GetProductsToOrderResponse []ProductInfo

func NewGetProductsToOrderResponse(res *orderpb.GetProductsToOrderResponse) (GetProductsToOrderResponse, error) {
	var response GetProductsToOrderResponse

	err := json.Unmarshal([]byte(res.Data), &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
