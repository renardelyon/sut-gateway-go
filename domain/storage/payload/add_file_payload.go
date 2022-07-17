package payload

import (
	"mime/multipart"
)

type AddFilePayload struct {
	File *multipart.FileHeader `form:"file" json:"file" binding:"required"`
}
