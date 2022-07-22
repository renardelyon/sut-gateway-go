package payload

// swagger:parameters DeleteToken
type SwaggerDeleteToken struct {
	// in: body
	Body DeleteTokenPayload
}

type DeleteTokenPayload struct {
	// example: jwtToken
	Token string `json:"token"`
}
