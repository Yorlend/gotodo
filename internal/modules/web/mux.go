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

	r.HandleFunc("/todos", th.Get).Methods("GET")
	r.HandleFunc("/todos/{id}", th.Delete).Methods("DELETE")
	r.HandleFunc("/todos", th.Post).Methods("POST")
	r.HandleFunc("/todos/{id}", th.Patch).Methods("PATCH")

	return r
}
