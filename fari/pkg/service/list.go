package service

import (
	"github.com/eserzhan/rest"
	"github.com/eserzhan/rest/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list todo.Todo_lists) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) Get(userId int) ([]todo.Todo_lists, error){
	return s.repo.Get(userId)
}

func (s *TodoListService) GetById(userId int, lstId string) (todo.Todo_lists, error){
	return s.repo.GetById(userId, lstId)
}

func (s *TodoListService) Delete(userId int, lstId string) (error){
	return s.repo.Delete(userId, lstId)
}

func (s *TodoListService) Update(userId int, lstId string, up todo.UpdateTodoLists) (error){
	return s.repo.Update(userId, lstId, up)
}
