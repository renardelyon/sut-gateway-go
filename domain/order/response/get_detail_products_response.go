package response

import orderpb "sut-gateway-go/pb/order"

type GetDetailProductsResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Url         string `json:"url"`
}

func NewGetDetailProductsResponse(res *orderpb.GetDetailProductsResponse) GetDetailProductsResponse {
	return GetDetailProductsResponse{
		Name:        res.Detailproduct.Name,
		Description: res.Detailproduct.Description,
		Image:       res.Detailproduct.Image,
		Url:         res.Detailproduct.Url,
	}
}
