package libs

import (
	"github.com/backent/ai-golang/controllers/controllers_auth"
	"github.com/backent/ai-golang/controllers/controllers_question"
	"github.com/backent/ai-golang/exceptions"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(authController controllers_auth.AuthControllerInterface, questionController controllers_question.QuestionControllerInterface) *httprouter.Router {
	router := httprouter.New()

	router.POST("/login", authController.Login)
	router.POST("/questions", questionController.Create)
	router.GET("/questions", questionController.GetAll)
	router.GET("/questions/:id", questionController.GetById)

	router.PanicHandler = exceptions.RouterPanicHandler

	return router
}
