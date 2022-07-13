package http_handler

import "sut-gateway-go/pb/auth"

type handler struct {
	auth auth.AuthServiceClient
}

type Handler interface {
}

func NewAuthHandler(auth auth.AuthServiceClient) Handler {
	return &handler{
		auth: auth,
	}
}
