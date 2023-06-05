package service

import (
	"github.com/eserzhan/rest"
	"github.com/eserzhan/rest/pkg/repository"
)

type TodoItemService struct {
	repo repository.TodoItem
}

func NewTodoItemService(repo repository.TodoItem) *TodoItemService {
	return &TodoItemService{repo: repo}
}

func (s *TodoItemService) Create(listId string, Item todo.Todo_items) (int, error) {
	return s.repo.Create(listId, Item)
}

func (s *TodoItemService) Get(userId int, lstId string) ([]todo.Todo_items, error){
	return s.repo.Get(userId, lstId)
}

func (s *TodoItemService) GetById(userId int, itemId string) (todo.Todo_items, error){
	return s.repo.GetById(userId, itemId)
}

func (s *TodoItemService) Delete(userId int, itemId string) (error){
	return s.repo.Delete(userId, itemId)
}

func (s *TodoItemService) Update(userId int, itemId string, up todo.UpdateTodoItems) (error){
	return s.repo.Update(userId, itemId, up)
}
