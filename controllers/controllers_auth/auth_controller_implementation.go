package controllers_auth

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/backent/ai-golang/helpers"
	"github.com/backent/ai-golang/middlewares"
	"github.com/backent/ai-golang/repositories/repositories_auth"
	"github.com/backent/ai-golang/services/services_auth"
	"github.com/backent/ai-golang/web"
	"github.com/backent/ai-golang/web/web_auth"
	"github.com/julienschmidt/httprouter"
)

type AuthControllerImplementation struct {
	services_auth.ServiceAuthInterface
	repositories_auth.RepositoryAuthInterface
}

func NewAuthControllerImplementation(auth services_auth.ServiceAuthInterface, repoAuth repositories_auth.RepositoryAuthInterface) AuthControllerInterface {
	return &AuthControllerImplementation{ServiceAuthInterface: auth, RepositoryAuthInterface: repoAuth}
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

func (implementation *AuthControllerImplementation) CurrentUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx := context.WithValue(r.Context(), helpers.ContextKey("token"), r.Header.Get("Authorization"))
	username := middlewares.ValidateToken(ctx, implementation.RepositoryAuthInterface)

	webResponse := web.WebResponse{
		Status: 200,
		Data: struct {
			Username string `json:"username"`
		}{
			Username: username,
		},
	}

	helpers.ReturnJSON(w, webResponse)
}
