package usecase

import (
	"github.com/linjiansi/TodoAPI/domain/model"
	"github.com/linjiansi/TodoAPI/domain/repository"
)

type TodoUsecase interface {
	Create(title, content string) (*model.Todo, error)
	FindBy(id int) (*model.Todo, error)
	Update(id int, title, content string) (*model.Todo, error)
	Delete(id int) (*model.Todo, error)
}

type todoUsecase struct {
	todoRepo repository.TodoRepository
}

func NewTodoUsecase(todoRepo repository.TodoRepository) TodoUsecase {
	return &todoUsecase{todoRepo: todoRepo}
}

func (tu *todoUsecase) Create(title, content string) (*model.Todo, error) {
	todo, err := model.NewTodo(title, content)
	if err != nil {
		return nil, err
	}

	createdTodo, err := tu.todoRepo.Create(todo)
	if err != nil {
		return nil, err
	}

	return createdTodo, nil
}

func (tu *todoUsecase) FindBy(id int) (*model.Todo, error) {
	foundTodo, err := tu.todoRepo.FindBy(id)
	if err != nil {
		return nil, err
	}

	return foundTodo, nil
}

func (tu *todoUsecase) Update(id int, title, content string) (*model.Todo, error) {
	targetTodo, err := tu.todoRepo.FindBy(id)
	if err != nil {
		return nil, err
	}

	err = targetTodo.Set(title, content)
	if err != nil {
		return nil, err
	}

	updatedTodo, err := tu.todoRepo.Update(targetTodo)
	if err != nil {
		return nil, err
	}

	return updatedTodo, nil
}

func (tu *todoUsecase) Delete(id int) (*model.Todo, error) {
	targetTodo, err := tu.todoRepo.FindBy(id)
	if err != nil {
		return nil, err
	}

	deletedTask, err := tu.todoRepo.Delete(targetTodo)
	if err != nil {
		return nil, err
	}

	return deletedTask, nil
}



