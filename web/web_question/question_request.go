package web_question

import "mime/multipart"

type QuestionPostRequest struct {
	Name       string `json:"name"`
	Amount     int    `json:"amount"`
	File       multipart.File
	FileHeader *multipart.FileHeader
}
