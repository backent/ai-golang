package exceptions

import (
	"net/http"

	"github.com/backent/ai-golang/helpers"
	"github.com/backent/ai-golang/web"
	"github.com/go-playground/validator/v10"
)

func RouterPanicHandler(w http.ResponseWriter, r *http.Request, i interface{}) {
	var response web.WebResponse
	// log.Printf("ERROR: %v\n%s", i, debug.Stack())

	if err, ok := i.(validator.ValidationErrors); ok {
		response = web.WebResponse{
			Status: http.StatusBadRequest,
			Data:   err.Error(),
		}
	} else if err, ok := i.(BadRequestError); ok {
		response = web.WebResponse{
			Status: http.StatusBadRequest,
			Data:   err.Error(),
		}
	} else if err, ok := i.(Unauthorized); ok {
		response = web.WebResponse{
			Status: http.StatusUnauthorized,
			Data:   err.Error(),
		}
	} else if err, ok := i.(Forbidden); ok {
		response = web.WebResponse{
			Status: http.StatusForbidden,
			Data:   err.Error(),
		}
	} else if err, ok := i.(NotFoundError); ok {
		response = web.WebResponse{
			Status: http.StatusNotFound,
			Data:   err.Error(),
		}
	} else if err, ok := i.(ConflictError); ok {
		response = web.WebResponse{
			Status: http.StatusConflict,
			Data:   err.Error(),
		}
	} else if err, ok := i.(error); ok {
		response = web.WebResponse{
			Status: http.StatusInternalServerError,
			Data:   err.Error(),
		}
	} else {
		response = web.WebResponse{
			Status: http.StatusInternalServerError,
			Data:   i,
		}
	}

	helpers.ReturnJSON(w, response)
}
