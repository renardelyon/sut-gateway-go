package response

// swagger:response GenerateNewTokenResponse
type SwaggerGenerateNewTokenResponse struct {
	// in: body
	Body GenerateNewTokenResponse
}

type GenerateNewTokenResponse struct {
	AccessToken string `json:"accessToken"`
}

func NewGenerateTokenResponse(token string) GenerateNewTokenResponse {
	return GenerateNewTokenResponse{
		AccessToken: token,
	}
}
