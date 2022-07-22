package response

import (
	authpb "sut-gateway-go/pb/auth"
)

// swagger:response LoginResponse
type SwaggerLoginResponse struct {
	// in: body
	Body struct {
		LoginResponse
	}
}

// swagger:response UserInfo
type SwaggerUserInfo struct {
	// in: body
	Body UserInfo
}

type UserInfo struct {
	Username string `json:"username"`
	Id       string `json:"id"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	AdminId  string `json:"adminId"`
}

type LoginResponse struct {
	AccessToken  string   `json:"accessToken"`
	RefreshToken string   `json:"refreshToken"`
	UserInfo     UserInfo `json:"userInfo"`
}

func NewLoginResponse(res *authpb.LoginResponse) LoginResponse {
	return LoginResponse{
		AccessToken:  res.Token,
		RefreshToken: res.Refreshtoken,
		UserInfo: UserInfo{
			Username: res.UserInfo.Username,
			Id:       res.UserInfo.Id,
			Name:     res.UserInfo.Name,
			Role:     res.UserInfo.Role.String(),
			AdminId:  res.UserInfo.AdminId,
		},
	}
}
