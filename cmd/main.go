package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/hello", ourBasicHandler)
	

	server := &http.Server{
		Addr: ":3333",
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Failed to listen to server", err)
	}
}

func ourBasicHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello, world!"))
}