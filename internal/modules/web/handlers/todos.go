package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yorlend/gotodo/internal/models"
	"go.uber.org/zap"
)

type TodosHandler struct {
	logger *zap.SugaredLogger
	repo   models.TodoRepository
}

func NewTodosHandler(
	logger *zap.SugaredLogger,
	repo models.TodoRepository,
) *TodosHandler {
	return &TodosHandler{
		logger: logger,
		repo:   repo,
	}
}

func (h *TodosHandler) Get(w http.ResponseWriter, r *http.Request) {
	todos, err := h.repo.ReadAll()
	if err != nil {
		h.logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// todos to json
	todosJson, err := json.Marshal(todos)
	if err != nil {
		h.logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(todosJson)
}

func (h *TodosHandler) Post(w http.ResponseWriter, r *http.Request) {
	todo := &models.Todo{}
	err := json.NewDecoder(r.Body).Decode(todo)
	if err != nil {
		h.logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.repo.Create(todo.Title, todo.Description)
	if err != nil {
		h.logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *TodosHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		h.logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.repo.Delete(id)
	if err != nil {
		h.logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *TodosHandler) Patch(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		h.logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	todo := &models.Todo{}
	err = json.NewDecoder(r.Body).Decode(todo)
	if err != nil {
		h.logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	todo.Id = id
	err = h.repo.Update(todo)
	if err != nil {
		h.logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
