package web_question

import "mime/multipart"

type QuestionPostRequest struct {
	Description string `json:"description"`
	File        multipart.File
	FileHeader  *multipart.FileHeader
}
