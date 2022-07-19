package response

import (
	"bytes"
	"encoding/json"
	productpb "sut-gateway-go/pb/product"

	"github.com/golang/protobuf/jsonpb"
)

type ProductInfo struct {
	Name string `json:"productName"`
	Url  string `json:"url"`
}

type GetAllProductsResponse struct {
	ProductInfos []ProductInfo `json:"products"`
}

func NewGetAllProducts(res *productpb.GetAllProductsResponse) (GetAllProductsResponse, error) {
	var response GetAllProductsResponse
	var buf bytes.Buffer

	err := (&jsonpb.Marshaler{}).Marshal(&buf, res)
	if err != nil {
		return response, err
	}

	jsonString := buf.String()

	err = json.Unmarshal([]byte(jsonString), &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
