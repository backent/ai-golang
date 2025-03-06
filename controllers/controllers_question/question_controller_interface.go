package controllers_question

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type QuestionControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	GetAll(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	GetById(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	DeleteById(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	CheckMaterial(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
