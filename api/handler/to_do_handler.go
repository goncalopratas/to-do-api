package handler

import (
	"api/domain"
	"api/usecase"
	"encoding/json"
	"net/http"
)

type ToDoHandler struct {
	ToDoUsecase *usecase.ToDoUsecase
}

func NewTodoHandler(tu *usecase.ToDoUsecase) *ToDoHandler {
	return &ToDoHandler{ToDoUsecase: tu}
}

func (h *ToDoHandler) GetAllToDos(w http.ResponseWriter, r *http.Request) {
	toDos, err := h.ToDoUsecase.GetTodos()

	if err != nil {
		http.Error(w, "Unable to fetch TODOS", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(toDos)

}

func (h *ToDoHandler) CreateToDo(w http.ResponseWriter, r *http.Request) {
	var toDo domain.ToDo
	err := json.NewDecoder(r.Body).Decode(&toDo)

	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = h.ToDoUsecase.CreateToDo(toDo)
	if err != nil {
		http.Error(w, "Unable to create TO DO", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(toDo)

}

func (h *ToDoHandler) DeleteToDo(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id", http.StatusBadRequest)
		return
	}

	err := h.ToDoUsecase.DeleteToDo(id)
	if err != nil {
		http.Error(w, "TO DO not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
