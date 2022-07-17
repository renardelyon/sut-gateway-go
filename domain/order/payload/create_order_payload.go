package payload

type CreateOrderPayload struct {
	Url         string `json:"url"`
	ProductName string `json:"productName"`
}
