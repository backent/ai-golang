package controllers_question

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type QuestionControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
