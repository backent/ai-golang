package controllers_exam

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/backent/ai-golang/helpers"
	"github.com/backent/ai-golang/middlewares"
	"github.com/backent/ai-golang/repositories/repositories_auth"
	"github.com/backent/ai-golang/services/services_exam"
	"github.com/backent/ai-golang/web"
	"github.com/backent/ai-golang/web/web_exam"
	"github.com/julienschmidt/httprouter"
)

type ExamControllerImplementation struct {
	services_exam.ExamServiceInterface
	repositories_auth.RepositoryAuthInterface
}

func NewExamControllerImplementation(exam services_exam.ExamServiceInterface, auth repositories_auth.RepositoryAuthInterface) ExamControllerInterface {
	return &ExamControllerImplementation{ExamServiceInterface: exam, RepositoryAuthInterface: auth}
}

func (implementation *ExamControllerImplementation) GetByQuestionId(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx := context.WithValue(r.Context(), helpers.ContextKey("token"), r.Header.Get("Authorization"))
	ctx = context.WithValue(ctx, helpers.ContextKey("username"), middlewares.ValidateToken(ctx, implementation.RepositoryAuthInterface))

	id, err := strconv.Atoi(p.ByName("id"))
	helpers.PanicIfError(err)
	data := implementation.ExamServiceInterface.GetExamByQuestionId(ctx, id)

	response := web.WebResponse{
		Status: 200,
		Data:   data,
	}

	helpers.ReturnJSON(w, response)

}
func (implementation *ExamControllerImplementation) Submit(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx := context.WithValue(r.Context(), helpers.ContextKey("token"), r.Header.Get("Authorization"))
	ctx = context.WithValue(ctx, helpers.ContextKey("username"), middlewares.ValidateToken(ctx, implementation.RepositoryAuthInterface))

	var request web_exam.ExamSubmitRequest

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&request)
	implementation.ExamServiceInterface.Submit(ctx, request)

	response := web.WebResponse{
		Status: 200,
	}

	helpers.ReturnJSON(w, response)

}
