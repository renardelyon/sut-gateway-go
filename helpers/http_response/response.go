package http_response

// swagger:response swaggerResponse
type SwaggerResponse struct {
	// in: body
	Body Response
}

type Response struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Status     string      `json:"status"`
	Timestamp  string      `json:"timestamp"`
	Data       interface{} `json:"data"`
}

type Empty struct{}
