package response

import notifpb "sut-gateway-go/pb/notification"

// swagger:response GetNotificationResponse
type SwaggerGetNotificationResponse struct {
	// in: body
	Body GetNotificationResponse
}

type GetNotificationResponse struct {
	Accepted int64 `json:"accepted"`
	Rejected int64 `json:"rejected"`
}

func NewGetNotificationResponse(res *notifpb.GetNotificationResponse) GetNotificationResponse {
	return GetNotificationResponse{
		Accepted: res.Accepted,
		Rejected: res.Rejected,
	}
}
