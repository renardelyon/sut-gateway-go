package payload

// swagger:parameters RegisterUserPayload
type SwaggerRegisterUserPayload struct {
	// in: body
	Body RegisterUserPayload
}

type RegisterUserPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	AdminId  string `json:"adminId"`
}
