package controllers_auth

import (
	"net/http"

	"github.com/backent/ai-golang/helpers"
	"github.com/backent/ai-golang/web"
	"github.com/julienschmidt/httprouter"
)

type AuthControllerImplementation struct {
}

func NewAuthControllerImplementation() AuthControllerInterface {
	return &AuthControllerImplementation{}
}

func (implementation *AuthControllerImplementation) Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	webResponse := web.WebResponse{
		Status: 200,
		Data:   "Login Succees",
	}

	helpers.ReturnJSON(w, webResponse)
}
