package api

import (
	"encoding/json"
	"net/http"
)

// TodoTypes
type Todo struct {
	id          string
	title       string
	description string
}

type TodoParams struct {
	id string
}

type TodoResponse struct {
	code int
	data Todo
}

type TodosResponse struct {
	code int
	data []Todo
}

// UserTypes
type User struct {
	id        string
	firstName string
	lastName  string
	username  string
}

type UserParams struct {
	id string
}

type UserResponse struct {
	code int
	data User
}

type UsersResponse struct {
	code int
	data []User
}

// Error Response
type Error struct {
	code    int
	message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		code:    code,
		message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An Unexpected Error Occurred.", http.StatusInternalServerError)
	}
)
