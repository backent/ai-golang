package web_question

import "mime/multipart"

type QuestionPostRequest struct {
	Name       string `json:"name"`
	Chapter    string `json:"chapter"`
	Amount     int    `json:"amount"`
	File       multipart.File
	FileHeader *multipart.FileHeader
}

type QuestionCheckFileMaterialRequest struct {
	Chapter    string `json:"chapter"`
	File       multipart.File
	FileHeader *multipart.FileHeader
}
