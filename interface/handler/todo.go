package handler

import (
	"github.com/linjiansi/TodoAPI/usecase"
	"net/http"
)

type TodoHandler interface {
	Post(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Put(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type todoHandler struct {
	todoUsecase usecase.TodoUsecase
}

func NewTodoHandler(todoUsecase usecase.TodoUsecase) TodoHandler {
	return &todoHandler{todoUsecase: todoUsecase}
}

func (th *todoHandler) Post(w http.ResponseWriter, r *http.Request) {
	
}

func (th *todoHandler) Get(w http.ResponseWriter, r *http.Request) {

}

func (th *todoHandler) Put(w http.ResponseWriter, r *http.Request) {

}

func (th *todoHandler) Delete(w http.ResponseWriter, r *http.Request) {

}
