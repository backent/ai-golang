package controllers_exam

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ExamControllerInterface interface {
	GetByQuestionId(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	GetById(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Submit(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
