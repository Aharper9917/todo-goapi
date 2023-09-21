package api

import (
	"encoding/json"
	"github.com/Aharper9917/todo-goapi/internal/tools"
	"net/http"
)

type TodoParams struct {
	Id string
}
type TodoResponse struct {
	Code int
	Data tools.TodoDetails
}
type TodosResponse struct {
	Dode int
	Data []tools.TodoDetails
}

type UserParams struct {
	Id string
}
type UserResponse struct {
	Code int
	Data tools.UserDetails
}
type UsersResponse struct {
	Code int
	Data []tools.UserDetails
}

type ErrorResponse struct {
	code    int
	message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := ErrorResponse{
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
