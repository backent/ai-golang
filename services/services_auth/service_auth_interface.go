package services_auth

import "github.com/backent/ai-golang/web/web_auth"

type ServiceAuthInterface interface {
	Login(request web_auth.AuthPostRequest) web_auth.AuthPostResponse
}
