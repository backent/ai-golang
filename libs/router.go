package libs

import (
	"github.com/backent/ai-golang/controllers/controllers_auth"
	"github.com/backent/ai-golang/controllers/controllers_exam"
	"github.com/backent/ai-golang/controllers/controllers_question"
	"github.com/backent/ai-golang/exceptions"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(
	authController controllers_auth.AuthControllerInterface,
	questionController controllers_question.QuestionControllerInterface,
	examController controllers_exam.ExamControllerInterface) *httprouter.Router {
	router := httprouter.New()

	router.POST("/login", authController.Login)
	router.POST("/questions", questionController.Create)
	router.POST("/question-check-material", questionController.CheckMaterial)
	router.GET("/current-user", authController.CurrentUser)
	router.GET("/questions", questionController.GetAll)
	router.GET("/questions/:id", questionController.GetById)
	router.DELETE("/questions/:id", questionController.DeleteById)
	router.GET("/exams/:id", examController.GetByQuestionId)
	router.GET("/exams-preview/:id", examController.GetById)
	router.POST("/exams", examController.Submit)

	router.PanicHandler = exceptions.RouterPanicHandler

	return router
}
