package payload

type RegisterAdminPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	AdminKey string `json:"adminKey"`
}
