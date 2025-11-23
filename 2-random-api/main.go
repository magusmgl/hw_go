package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

type CustomHandler struct{}

func NewCustomHandler(router *http.ServeMux) {
	mux := &CustomHandler{}
	router.HandleFunc("/randInt", mux.RandomNum)
}

func (handler *CustomHandler) RandomNum(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL.Path)
	num := rand.Intn(7)
	numStr := strconv.Itoa(num)
	_, err := w.Write([]byte(numStr))
	if err != nil {
		http.Error(w, "Не удалось записать ответ", http.StatusInternalServerError)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	mux := http.NewServeMux()
	NewCustomHandler(mux)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Fprintln(os.Stdout, "Запущен сервер на протоколе 8080")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Server error")
	}
}
