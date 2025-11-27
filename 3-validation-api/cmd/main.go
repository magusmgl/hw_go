package main

import (
	"fmt"
	"net/http"
	"validation/api/configs"
	"validation/api/internal/valid"
)

func main() {
	conf := configs.LoadConfig()

	router := http.NewServeMux()
	valid.NewEmailHandler(router, valid.EmailHandlerDeps{
		Config: conf,
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Run server on port 8080")
	server.ListenAndServe()
}
