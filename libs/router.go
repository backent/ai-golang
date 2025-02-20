package libs

import (
	"net/http"

	"github.com/backent/ai-golang/controllers/controllers_auth"
	"github.com/backent/ai-golang/controllers/controllers_question"
	"github.com/backent/ai-golang/helpers"
	"github.com/backent/ai-golang/web"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(authController controllers_auth.AuthControllerInterface, questionController controllers_question.QuestionControllerInterface) *httprouter.Router {
	router := httprouter.New()

	router.POST("/login", authController.Login)
	router.POST("/questions", questionController.Create)

	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}) {
		response := web.WebResponse{
			Status: 500,
			Data:   i,
		}
		helpers.ReturnJSON(w, response)
	}

	return router
}
