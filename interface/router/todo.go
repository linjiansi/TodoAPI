package router

import (
	"github.com/linjiansi/TodoAPI/interface/handler"
	"net/http"
)

type TodoRouter interface {
	HandleTodosRequest(w http.ResponseWriter, r *http.Request)
}

type todoRouter struct {
	todoHandler handler.TodoHandler
}

func NewTodoRouter(todoHandler handler.TodoHandler) TodoRouter {
	return &todoRouter{todoHandler: todoHandler}
}

func (tr todoRouter) HandleTodosRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tr.todoHandler.Get(w, r)
	case "POST":
		tr.todoHandler.Post(w, r)
	case "PUT":
		tr.todoHandler.Put(w, r)
	case "DELETE":
		tr.todoHandler.Delete(w, r)
	}
}
