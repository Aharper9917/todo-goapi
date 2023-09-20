package main

import (
	"fmt"
	"net/http"

	"github.com/Aharper9917/todo-goapi/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting TODO GO API service...")

	fmt.Println(`
  __________  ____  ____     __________     ___    ____  ____
 /_  __/ __ \/ __ \/ __ \   / ____/ __ \   /   |  / __ \/  _/
  / / / / / / / / / / / /  / / __/ / / /  / /| | / /_/ // /  
 / / / /_/ / /_/ / /_/ /  / /_/ / /_/ /  / ___ |/ ____// /   
/_/  \____/_____/\____/   \____/\____/  /_/  |_/_/   /___/   
                                                             `)

	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}

}
