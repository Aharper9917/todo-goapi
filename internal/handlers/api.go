package handlers

import (
	"fmt"
	"github.com/Aharper9917/todo-goapi/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	// Global middleware
	r.Use(chimiddle.StripSlashes)

	r.Route("/todos", func(router chi.Router) {
		//router.Use(middleware.Authorization)
		fmt.Println("todo router")

		router.Get("/", GetTodos)
		router.Post("/", CreateTodo)

		router.Route("/{todoID}", func(router chi.Router) {
			router.Use(TodoCtx)
			router.Get("/", GetTodoById)
			router.Put("/", UpdateTodoById)
			router.Delete("/", DeleteTodoById)
		})
	})

	r.Route("/users", func(router chi.Router) {
		router.Use(middleware.Authorization)

		router.Get("/", GetUsers)
		router.Post("/", CreateUser)

		router.Route("/{userID}", func(router chi.Router) {
			router.Get("/", GetUserById)
			router.Put("/", UpdateUserById)
			router.Delete("/", DeleteUserById)
		})
	})
}
