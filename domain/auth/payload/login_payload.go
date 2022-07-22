package payload

// swagger:parameters LoginPayload
type SwaggerLoginPayload struct {
	// in: body
	Body LoginPayload
}

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
