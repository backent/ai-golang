package controllers_auth

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AuthControllerInterface interface {
	Login(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
