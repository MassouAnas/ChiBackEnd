package application

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/MassouAnas/ChiBackEnd/handler"
)

func listRoutes() *chi.Mux{
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter,r *http.Request){
		w.WriteHeader(http.StatusOK)
	})
	router.Route("/Todo", LoadTodoRoutes)
	return router
}

func LoadTodoRoutes(router chi.Router ){
	todoHandler := &handler.Todo{}

	router.Post("/", todoHandler.Create)
	router.Get("/", todoHandler.List)
	router.Get("/{id}", todoHandler.ListByID)
	router.Put("/{id}", todoHandler.Update)
	router.Delete("/{id}", todoHandler.DeleteById)
}