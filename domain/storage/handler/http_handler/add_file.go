package http_handler

import (
	"bufio"
	"context"
	"io"
	"net/http"
	"sut-gateway-go/domain/storage/payload"
	"sut-gateway-go/helpers/http_response"
	"sut-gateway-go/helpers/timestamp"
	"sut-gateway-go/pb/auth"
	storagepb "sut-gateway-go/pb/storage"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *handler) AddFile(ctx *gin.Context) {
	select {
	case <-ctx.Done():
		return
	case <-ctx.Request.Cancel:
		return
	default:
	}

	userInfo, _ := ctx.Get("userInfo")

	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := h.storage.AddFile(c)
	if err != nil {
		response := http_response.Response{
			StatusCode: http.StatusBadGateway,
			Message:    err.Error(),
			Status:     "Bad Gateway",
			Timestamp:  timestamp.GetNow(),
		}
		ctx.JSON(http.StatusBadGateway, response)
		return
	}

	req := &storagepb.UploadRequest{
		Data: &storagepb.UploadRequest_Info{
			Info: &storagepb.FileInfo{
				UserId: userInfo.(*auth.UserInfo).Id,
			},
		},
	}

	err = stream.Send(req)
	if err != nil {
		response := http_response.Response{
			StatusCode: http.StatusBadGateway,
			Message:    err.Error(),
			Status:     "Bad Gateway",
			Timestamp:  timestamp.GetNow(),
		}
		ctx.JSON(http.StatusBadGateway, response)
		return
	}

	reqPayload := payload.AddFilePayload{}
	if err := ctx.ShouldBind(&reqPayload); err != nil {
		response := http_response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Status:     "Bad Gateway",
			Timestamp:  timestamp.GetNow(),
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	file, err := reqPayload.File.Open()
	if err != nil {
		response := http_response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Status:     "Bad Gateway",
			Timestamp:  timestamp.GetNow(),
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}

		if err != nil {
			response := http_response.Response{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Status:     "Bad Gateway",
				Timestamp:  timestamp.GetNow(),
			}
			ctx.JSON(http.StatusInternalServerError, response)
			return
		}

		req := &storagepb.UploadRequest{
			Data: &storagepb.UploadRequest_ChunkData{
				ChunkData: buffer[:n],
			},
		}

		err = stream.Send(req)
		if err != nil {
			response := http_response.Response{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Status:     "Bad Gateway",
				Timestamp:  timestamp.GetNow(),
			}
			ctx.JSON(http.StatusInternalServerError, response)
			return
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		response := http_response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Status:     "Internal Server Error",
			Timestamp:  timestamp.GetNow(),
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := http_response.Response{
		StatusCode: http.StatusOK,
		Message:    "File successfully uploaded",
		Status:     "OK",
		Timestamp:  timestamp.GetNow(),
		Data:       res.Id,
	}

	ctx.JSON(http.StatusOK, response)
}
