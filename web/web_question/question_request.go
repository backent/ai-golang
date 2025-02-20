package web_question

import "mime/multipart"

type QuestionPostRequest struct {
	Amount     int `json:"amount"`
	File       multipart.File
	FileHeader *multipart.FileHeader
}
