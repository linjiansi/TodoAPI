package infrastructure

import (
	"database/sql"
	"github.com/linjiansi/TodoAPI/domain/model"
	"github.com/linjiansi/TodoAPI/domain/repository"
)

type TodoRepository struct {
	*sql.DB
}

func NewUserReposioty(db *sql.DB) repository.TodoRepository {
	return &TodoRepository{db}
}

func (r *TodoRepository) Create(todo *model.Todo) (*model.Todo, error) {
	stmt, err := r.Prepare(`INSERT INTO todos(title, content) VALUES (?, ?)`)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	if _, err := stmt.Exec(todo.Title, todo.Content); err != nil {
		return nil, err
	}

	return todo, nil
}

func (r *TodoRepository) FindBy(id int) (*model.Todo, error) {
	stmt, err := r.Prepare(`SELECT id, title, content FROM todos WHERE id = ?`)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	todo := &model.Todo{}
	err = stmt.QueryRow(id).Scan(&todo.ID, &todo.Title, &todo.Content)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (r *TodoRepository) Update(todo *model.Todo) (*model.Todo, error) {
	stmt, err := r.Prepare(`UPDATE todos SET title = ?, content = ?. WHERE id = ?`)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	if _, err := stmt.Exec(todo.Title, todo.Content, todo.ID); err != nil {
		return nil, err
	}

	return todo, nil
}

func (r *TodoRepository) Delete(todo *model.Todo) (*model.Todo, error) {
	stmt, err := r.Prepare(`DELETE FROM todos WHERE id = ?`)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	if _, err := stmt.Exec(todo.ID); err != nil {
		return nil, err
	}

	return todo, nil
}