package payload

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
