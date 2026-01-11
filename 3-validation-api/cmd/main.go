package main

import (
	"fmt"
	"log"
	"net/http"
	"validation/api/configs"
	"validation/api/internal/valid"
)

func main() {
	conf, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	router := http.NewServeMux()
	valid.NewEmailHandler(router, valid.EmailHandlerDeps{
		Config: conf,
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Run server on port 8080")
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
