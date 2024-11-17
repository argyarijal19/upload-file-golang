package model

import "mime/multipart"

type UploadFile struct {
	FileName string                `json:"file_name" form:"file_name"`
	File     *multipart.FileHeader `json:"attachment"` // We will populate this manually
}

type Request struct {
	FileName string `json:"file_path"`
}
