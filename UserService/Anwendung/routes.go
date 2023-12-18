package Anwendung

import (
    "net/http"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/mmartin04/Docker/UserService/handler"
)

func loadRoutes() *chi.Mux {
    router := chi.NewRouter()
    router.Use(middleware.Logger)
    userHandler := &handler.User{}

    router.Route("/users", func(r chi.Router) {
        router.Post("/", userHandler.Create)
        router.Get("/{id}", userHandler.GetByID)       
        router.Get("/", orderHandler.List)
        router.Put("/{id}", orderHandler.UpdateByID)
        router.Delete("/{id}", orderHandler.DeleteByID)

    })

    return router
}
