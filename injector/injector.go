//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/backent/ai-golang/controllers/controllers_auth"
	"github.com/backent/ai-golang/controllers/controllers_exam"
	"github.com/backent/ai-golang/controllers/controllers_question"
	"github.com/backent/ai-golang/libs"
	"github.com/backent/ai-golang/repositories/repositories_auth"
	"github.com/backent/ai-golang/repositories/repositories_exam"
	"github.com/backent/ai-golang/repositories/repositories_question"
	"github.com/backent/ai-golang/repositories/repositories_storage"
	"github.com/backent/ai-golang/services/services_ai"
	"github.com/backent/ai-golang/services/services_auth"
	"github.com/backent/ai-golang/services/services_exam"
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
	repositories_question.NewRepositoryQuestionImplementation,
)

var ExamSet = wire.NewSet(
	controllers_exam.NewExamControllerImplementation,
	services_exam.NewExamServiceImplementation,
	repositories_exam.NewExamRepositoryImplementation,
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
		libs.NewDatabase,
		RepositoriesSet,
		AuthSet,
		QuestionSet,
		ExamSet,
		AiSet,
	)

	return nil
}
