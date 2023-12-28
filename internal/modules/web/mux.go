package web

import (
	"github.com/gorilla/mux"
	"github.com/yorlend/gotodo/internal/modules/web/handlers"
)

func NewRouter(
	hh *handlers.HomeHandler,
	th *handlers.TodosHandler,
) *mux.Router {
	r := mux.NewRouter()

	r.Handle("/", hh).Methods("GET")

	r.HandleFunc("/api/todos", th.Get).Methods("GET")
	r.HandleFunc("/api/todos/{id}", th.Delete).Methods("DELETE")
	r.HandleFunc("/api/todos", th.Post).Methods("POST")
	r.HandleFunc("/api/todos/{id}", th.Patch).Methods("PATCH")
	r.HandleFunc("/api/todos/{id}", th.GetOne).Methods("GET")

	return r
}
