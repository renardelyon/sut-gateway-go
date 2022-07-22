package payload

// swagger:parameters GenerateNewTokenPayload
type SwaggerGenerateNewTokenPayload struct {
	// in: body
	Body GenerateNewTokenPayload
}

type GenerateNewTokenPayload struct {
	Token string `json:"token"`
}
