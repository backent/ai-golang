package controllers_question

import (
	"net/http"

	"github.com/backent/ai-golang/helpers"
	"github.com/backent/ai-golang/services/services_question"
	"github.com/backent/ai-golang/web"
	"github.com/backent/ai-golang/web/web_question"
	"github.com/julienschmidt/httprouter"
)

type QuestionControllerImplementation struct {
	services_question.QuestionServiceInterface
}

func NewQuestionControllerImplementation(question services_question.QuestionServiceInterface) QuestionControllerInterface {
	return &QuestionControllerImplementation{
		QuestionServiceInterface: question,
	}
}

func (implementation *QuestionControllerImplementation) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	const MAX_FILE = 10 << 20 // 20MB
	err := r.ParseMultipartForm(MAX_FILE)
	helpers.PanicIfError(err)

	request := web_question.QuestionPostRequest{}

	file, fileHeader, err := r.FormFile("file")
	helpers.PanicIfError(err)

	request.Description = r.FormValue("description")
	request.File = file
	request.FileHeader = fileHeader

	data := implementation.QuestionServiceInterface.Create(request)

	webResponse := web.WebResponse{
		Status: 200,
		Data:   data,
	}

	helpers.ReturnJSON(w, webResponse)
}
