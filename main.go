package main

import (
	"net/http"

	_ "github.com/backent/ai-golang/config"
	"github.com/backent/ai-golang/injector"
)

func main() {
	router := injector.InitializeRouter()

	server := http.Server{
		Addr:    ":8022",
		Handler: router,
	}

	server.ListenAndServe()
}
