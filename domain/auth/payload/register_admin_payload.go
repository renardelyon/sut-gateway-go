package payload

// swagger:parameters RegisterAdminPayload
type SwaggerRegisterAdminPayload struct {
	// in:body
	Body RegisterAdminPayload
}
type RegisterAdminPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	AdminKey string `json:"adminKey"`
}
