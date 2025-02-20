//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/backent/ai-golang/controllers/controllers_auth"
	"github.com/backent/ai-golang/controllers/controllers_question"
	"github.com/backent/ai-golang/libs"
	"github.com/backent/ai-golang/repositories/repositories_auth"
	"github.com/backent/ai-golang/repositories/repositories_storage"
	"github.com/backent/ai-golang/services/services_ai"
	"github.com/backent/ai-golang/services/services_auth"
	"github.com/backent/ai-golang/services/services_question"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var AuthSet = wire.NewSet(
	controllers_auth.NewAuthControllerImplementation,
	services_auth.NewServiceAuthImplementation,
	repositories_auth.NewRepositoryAuthJWTImpl,
)

var QuestionSet = wire.NewSet(
	controllers_question.NewQuestionControllerImplementation,
	services_question.NewQuestionServiceImplementation,
)

var AiSet = wire.NewSet(
	services_ai.NewAiServiceGemini,
)

var RepositoriesSet = wire.NewSet(
	repositories_storage.NewStorageServiceLocalImplementation,
)

func InitializeRouter() *httprouter.Router {
	wire.Build(
		libs.NewRouter,
		RepositoriesSet,
		AuthSet,
		QuestionSet,
		AiSet,
	)

	return nil
}
