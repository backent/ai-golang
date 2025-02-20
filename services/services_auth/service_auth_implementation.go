package services_auth

import (
	"strings"

	"github.com/backent/ai-golang/helpers"
	"github.com/backent/ai-golang/repositories/repositories_auth"
	"github.com/backent/ai-golang/web/web_auth"
)

type ServiceAuthImplementation struct {
	repositories_auth.RepositoryAuthInterface
}

func NewServiceAuthImplementation(auth repositories_auth.RepositoryAuthInterface) ServiceAuthInterface {
	return &ServiceAuthImplementation{RepositoryAuthInterface: auth}
}

func (implementation *ServiceAuthImplementation) Login(request web_auth.AuthPostRequest) web_auth.AuthPostResponse {

	if !checkUserExists(request.Username) {
		return web_auth.AuthPostResponse{}
	}

	token, err := implementation.RepositoryAuthInterface.Issue(request.Username)
	helpers.PanicIfError(err)

	return web_auth.AuthPostResponse{Token: token}
}

func checkUserExists(username string) bool {
	availableUsername := make(map[string]interface{})
	availableUsername["teacher"] = nil
	availableUsername["student"] = nil

	_, ok := availableUsername[strings.ToLower(username)]
	return ok
}
