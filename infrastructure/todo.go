package infrastructure

import (
	"database/sql"
	"github.com/linjiansi/TodoAPI/domain/repository"
	"github.com/linjiansi/TodoAPI/domain/model"
)

type TodoRepository struct {
	*sql.DB
}

func NewUserReposioty(db *sql.DB) repository.TodoRepository {
	return &TodoRepository{db}
}

func (r *TodoRepository) Create(todo *model.Todo) (*model.Todo, error) {

}

func (r *TodoRepository) FindBy(id int) (*model.Todo, error) {

}

func (r *TodoRepository) Update(todo *model.Todo) (*model.Todo, error) {

}

func (r *TodoRepository) Delete(todo *model.Todo) (*model.Todo, error) {

}