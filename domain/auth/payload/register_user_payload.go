package payload

type RegisterUserPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	AdminId  string `json:"adminId"`
}
