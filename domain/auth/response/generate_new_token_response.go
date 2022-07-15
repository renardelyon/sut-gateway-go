package response

type GenerateNewTokenResponse struct {
	AccessToken string `json:"accessToken"`
}

func NewGenerateTokenResponse(token string) GenerateNewTokenResponse {
	return GenerateNewTokenResponse{
		AccessToken: token,
	}
}
