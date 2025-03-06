package controllers_question

import (
	"context"
	"net/http"
	"strconv"

	"github.com/backent/ai-golang/helpers"
	"github.com/backent/ai-golang/middlewares"
	"github.com/backent/ai-golang/repositories/repositories_auth"
	"github.com/backent/ai-golang/services/services_question"
	"github.com/backent/ai-golang/web"
	"github.com/backent/ai-golang/web/web_question"
	"github.com/julienschmidt/httprouter"
)

type QuestionControllerImplementation struct {
	services_question.QuestionServiceInterface
	repositories_auth.RepositoryAuthInterface
}

func NewQuestionControllerImplementation(question services_question.QuestionServiceInterface, auth repositories_auth.RepositoryAuthInterface) QuestionControllerInterface {
	return &QuestionControllerImplementation{
		QuestionServiceInterface: question,
		RepositoryAuthInterface:  auth,
	}
}

func (implementation *QuestionControllerImplementation) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx := context.WithValue(r.Context(), helpers.ContextKey("token"), r.Header.Get("Authorization"))
	ctx = context.WithValue(ctx, helpers.ContextKey("username"), middlewares.ValidateToken(ctx, implementation.RepositoryAuthInterface))

	const MAX_FILE = 10 << 20 // 20MB
	err := r.ParseMultipartForm(MAX_FILE)
	helpers.PanicIfError(err)

	request := web_question.QuestionPostRequest{}

	file, fileHeader, err := r.FormFile("file")
	helpers.PanicIfError(err)

	request.Name = r.FormValue("name")
	request.Chapter = r.FormValue("chapter")
	request.Amount, err = strconv.Atoi(r.FormValue("amount"))
	helpers.PanicIfError(err)
	request.File = file
	request.FileHeader = fileHeader

	data := implementation.QuestionServiceInterface.Create(ctx, request)

	webResponse := web.WebResponse{
		Status: 200,
		Data:   data,
	}

	helpers.ReturnJSON(w, webResponse)
}

func (implementation *QuestionControllerImplementation) GetAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx := context.WithValue(r.Context(), helpers.ContextKey("token"), r.Header.Get("Authorization"))
	ctx = context.WithValue(ctx, helpers.ContextKey("username"), middlewares.ValidateToken(ctx, implementation.RepositoryAuthInterface))

	data := implementation.QuestionServiceInterface.GetAll(ctx)

	webResponse := web.WebResponse{
		Status: 200,
		Data:   data,
	}

	helpers.ReturnJSON(w, webResponse)
}

func (implementation *QuestionControllerImplementation) GetById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx := context.WithValue(r.Context(), helpers.ContextKey("token"), r.Header.Get("Authorization"))
	ctx = context.WithValue(ctx, helpers.ContextKey("username"), middlewares.ValidateToken(ctx, implementation.RepositoryAuthInterface))

	id, err := strconv.Atoi(p.ByName("id"))
	helpers.PanicIfError(err)

	data := implementation.QuestionServiceInterface.GetById(ctx, id)

	webResponse := web.WebResponse{
		Status: 200,
		Data:   data,
	}

	helpers.ReturnJSON(w, webResponse)
}

func (implementation *QuestionControllerImplementation) DeleteById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx := context.WithValue(r.Context(), helpers.ContextKey("token"), r.Header.Get("Authorization"))
	ctx = context.WithValue(ctx, helpers.ContextKey("username"), middlewares.ValidateToken(ctx, implementation.RepositoryAuthInterface))

	id, err := strconv.Atoi(p.ByName("id"))
	helpers.PanicIfError(err)

	implementation.QuestionServiceInterface.DeleteById(ctx, id)

	webResponse := web.WebResponse{
		Status: 200,
		Data:   nil,
	}

	helpers.ReturnJSON(w, webResponse)
}

func (implementation *QuestionControllerImplementation) CheckMaterial(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx := context.WithValue(r.Context(), helpers.ContextKey("token"), r.Header.Get("Authorization"))
	ctx = context.WithValue(ctx, helpers.ContextKey("username"), middlewares.ValidateToken(ctx, implementation.RepositoryAuthInterface))

	const MAX_FILE = 10 << 20 // 20MB
	err := r.ParseMultipartForm(MAX_FILE)
	helpers.PanicIfError(err)

	request := web_question.QuestionCheckFileMaterialRequest{}

	file, fileHeader, err := r.FormFile("file")
	helpers.PanicIfError(err)

	request.Chapter = r.FormValue("chapter")
	helpers.PanicIfError(err)
	request.File = file
	request.FileHeader = fileHeader

	data := implementation.QuestionServiceInterface.CheckMaterial(ctx, request)

	webResponse := web.WebResponse{
		Status: 200,
		Data:   data,
	}

	helpers.ReturnJSON(w, webResponse)
}
