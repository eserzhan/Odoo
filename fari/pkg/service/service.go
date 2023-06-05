package service

import (
	"github.com/eserzhan/rest/pkg/repository"
	"github.com/eserzhan/rest"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list todo.Todo_lists) (int, error)
	Get(userId int) ([]todo.Todo_lists, error)
	GetById(userId int, lstId string) (todo.Todo_lists, error)
	Delete(userId int, lstId string) (error)
	Update(userId int, lstId string, up todo.UpdateTodoLists) (error)
}

type TodoItem interface {
	Create(lstId string, list todo.Todo_items) (int, error)
	Get(userId int, lstId string) ([]todo.Todo_items, error)
	GetById(userId int, itemId string) (todo.Todo_items, error)
	Delete(userId int, itemId string) (error)
	Update(userId int, itemId string, up todo.UpdateTodoItems) (error)
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		TodoList: NewTodoListService(repo.TodoList),
		TodoItem: NewTodoItemService(repo.TodoItem),
	}
}