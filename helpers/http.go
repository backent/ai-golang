package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/backent/ai-golang/web"
)

func ReturnJSON(w http.ResponseWriter, response web.WebResponse) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(response.Status)

	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}
