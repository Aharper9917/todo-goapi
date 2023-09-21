package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Aharper9917/todo-goapi/api"
	"github.com/Aharper9917/todo-goapi/internal/tools"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func TodoCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		todoID := chi.URLParam(r, "todoID")
		fmt.Printf("%+v\n", todoID)

		var db *tools.DatabaseInterface
		db, err := tools.NewDatabase()
		todo, err := (*db).GetTodoDetails(todoID)

		if err != nil || todo == nil {
			log.Error(err)
			api.InternalErrorHandler(w)
			return
		}

		ctx := context.WithValue(r.Context(), "todo", todo)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetTodos(w http.ResponseWriter, r *http.Request) {}

func GetTodoById(w http.ResponseWriter, r *http.Request) {
	var err error
	ctx := r.Context()
	fmt.Printf("%+v\n", ctx)
	todo, ok := ctx.Value("todo").(*tools.TodoDetails)
	if !ok {
		api.InternalErrorHandler(w)
		return
	}

	res := api.TodoResponse{
		Code: http.StatusOK,
		Data: *todo,
	}
	fmt.Printf("%+v\n", res)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(res)

	if err != nil {
		api.InternalErrorHandler(w)
		return
	}
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {}

func UpdateTodoById(w http.ResponseWriter, r *http.Request) {}

func DeleteTodoById(w http.ResponseWriter, r *http.Request) {}

//func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
//	var params = api.CoinBalanceParams{}
//	var decoder *schema.Decoder = schema.NewDecoder()
//	var err error
//
//	err = decoder.Decode(&params, r.URL.Query())
//
//	if err != nil {
//		log.ErrorResponse(err)
//		api.InternalErrorHandler(w)
//		return
//	}
//
//	var database *tools.DatabaseInterface
//	database, err = tools.NewDatabase()
//	if err != nil {
//		api.InternalErrorHandler(w)
//		return
//	}
//
//	var tokenDetails *tools.CoinDetails
//	tokenDetails = (*database).GetUserCoins(params.Username)
//	if tokenDetails == nil {
//		log.ErrorResponse(err)
//		api.InternalErrorHandler(w)
//		return
//	}
//
//	var response = api.CoinBalanceResponse{
//		Balance: (*tokenDetails).Coins,
//		Code:    http.StatusOK,
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	err = json.NewEncoder(w).Encode(response)
//	if err != nil {
//		log.ErrorResponse(err)
//		api.InternalErrorHandler(w)
//		return
//	}
//}
