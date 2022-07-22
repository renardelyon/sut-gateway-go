package payload

import (
	"mime/multipart"
)

// Excel File
// swagger:parameters AddFilePayload
type AddFilePayload struct {
	// Required:true
	// in: formData
	// swagger:file
	File *multipart.FileHeader `form:"file" json:"file" binding:"required"`
}
