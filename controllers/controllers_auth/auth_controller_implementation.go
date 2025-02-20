package controllers_auth

import (
	"encoding/json"
	"net/http"

	"github.com/backent/ai-golang/helpers"
	"github.com/backent/ai-golang/services/services_auth"
	"github.com/backent/ai-golang/web"
	"github.com/backent/ai-golang/web/web_auth"
	"github.com/julienschmidt/httprouter"
)

type AuthControllerImplementation struct {
	services_auth.ServiceAuthInterface
}

func NewAuthControllerImplementation(auth services_auth.ServiceAuthInterface) AuthControllerInterface {
	return &AuthControllerImplementation{ServiceAuthInterface: auth}
}

func (implementation *AuthControllerImplementation) Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var request web_auth.AuthPostRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	helpers.PanicIfError(err)

	tokenResponse := implementation.ServiceAuthInterface.Login(request)

	webResponse := web.WebResponse{
		Status: 200,
		Data:   tokenResponse,
	}

	helpers.ReturnJSON(w, webResponse)
}
