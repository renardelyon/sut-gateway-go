package http_handler

import (
	"context"
	"net/http"
	"sut-gateway-go/helpers/http_response"
	"sut-gateway-go/helpers/timestamp"
	"sut-gateway-go/pb/auth"

	"github.com/gin-gonic/gin"
)

type SwaggerUserAdmin struct {
	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Role     string `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
}

// swagger:response GetUserByRoleResponse
type SwaggerGetUserByRoleResponse struct {
	// in:body
	Body []SwaggerUserAdmin
}

// swagger:parameters GetUserByRoleQuery
type SwaggerGetUserByRoleQuery struct {
	// in: query
	// example: admin
	Role string
}

func (h *handler) GetUserByRole(ctx *gin.Context) {
	query, ok := ctx.GetQuery("role")

	if !ok {
		response := http_response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Query string role doesn't exist",
			Status:     "Bad request",
			Timestamp:  timestamp.GetNow(),
		}
		ctx.JSON(http.StatusUnauthorized, response)
		return
	}

	userInfo, _ := ctx.Get("userInfo")

	if userInfo.(*auth.UserInfo).Role != auth.Role_ADMIN {
		response := http_response.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "Unauthorized",
			Status:     "Unauthorized",
			Timestamp:  timestamp.GetNow(),
		}
		ctx.JSON(http.StatusUnauthorized, response)
		return
	}

	res, _ := h.auth.GetUserByRole(context.Background(), &auth.GetUserByRoleRequest{
		Role: query,
	})
	if res.Error != "" {
		response := http_response.Response{
			StatusCode: http.StatusBadGateway,
			Message:    res.Error,
			Status:     "Bad Gateway",
			Timestamp:  timestamp.GetNow(),
		}
		ctx.JSON(http.StatusBadGateway, response)
		return
	}

	response := http_response.Response{
		StatusCode: http.StatusOK,
		Message:    "OK",
		Status:     "OK",
		Timestamp:  timestamp.GetNow(),
		Data:       res.UserAdminInfo,
	}

	ctx.JSON(http.StatusOK, response)
}
