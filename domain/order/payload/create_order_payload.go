package payload

// swagger:parameters CreateOrderPayload
type SwaggerCreateOrderPayload struct {
	// in: body
	Body CreateOrderPayload
}

type CreateOrderPayload struct {
	Url         string `json:"url"`
	ProductName string `json:"productName"`
}
