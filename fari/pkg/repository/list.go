package repository

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/eserzhan/rest"
	"github.com/jmoiron/sqlx"
)

const (
	todoListTable = "todo_lists"

	usersListTable = "users_lists"
)

type TodoListRepository struct {
	db *sqlx.DB
}

func NewTodoListRepository(db *sqlx.DB) *TodoListRepository {
	return &TodoListRepository{db: db}
}

func (r *TodoListRepository) Create(userId int, list todo.Todo_lists) (int, error){
	
	tx, err := r.db.Begin()

	if err != nil {
		return 0, err 
	}
	defer tx.Rollback()

	var listId int 
	query_1 := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListTable)
	err = tx.QueryRow(query_1, list.Title, list.Description).Scan(&listId)

	if err != nil {
		return 0, err 
	}
	query_2 := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListTable)
	_, err = tx.Exec(query_2, userId, listId)

	if err != nil {
		return 0, err
	}

	return listId, tx.Commit() 
}

func (r *TodoListRepository) Get(userId int) ([]todo.Todo_lists, error) {
	id := strconv.Itoa(userId)
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl JOIN %s ul  ON tl.id = ul.list_id WHERE ul.user_id = $1", todoListTable, usersListTable)
	//query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl JOIN %s ul USING (list_id) WHERE ul.user_id = %s", todoListTable, usersListTable, id)

	var res []todo.Todo_lists
	err := r.db.Select(&res, query, id)

	if err != nil {
		return []todo.Todo_lists{}, err
	}

	return res, nil
}

func (r *TodoListRepository) GetById(userId int, lstId string) (todo.Todo_lists, error) {
	usId := strconv.Itoa(userId)

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl JOIN %s ul  ON tl.id = ul.list_id WHERE ul.user_id = $1 and tl.id = $2", todoListTable, usersListTable)
	//query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl JOIN %s ul USING (list_id) WHERE ul.user_id = %s", todoListTable, usersListTable, id)

	var res todo.Todo_lists
	err := r.db.Get(&res, query, usId, lstId)

	if err != nil {
		return todo.Todo_lists{}, err
	}

	return res, nil
}

func (r *TodoListRepository) Delete(userId int, lstId string) (error) {
	usId := strconv.Itoa(userId)

	//query := fmt.Sprintf("DELETE FROM %s ul USING %s tl WHERE tl.id = ul.list_id and ul.user_id = $1 and ul.list_id = $2", todoListTable, usersListTable)
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id=$1 AND ul.list_id=$2",
	todoListTable, usersListTable)


	_, err := r.db.Exec(query, usId, lstId)


	return err
}


func (r *TodoListRepository) Update(userId int, lstId string, up todo.UpdateTodoLists) (error) {
	usId := strconv.Itoa(userId)

	fields := make([]string, 0)
	placeholders := make([]interface{}, 0)
	args := 1

	if up.Description != nil {
		fields = append(fields, fmt.Sprintf("description=$%d", args))
		placeholders = append(placeholders, *up.Description)
		args += 1
	}

	if up.Title != nil {
		fields = append(fields, fmt.Sprintf("title=$%d", args))
		placeholders = append(placeholders, *up.Title)
		args += 1
	}
	
	changes := strings.Join(fields, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.user_id = $%d AND ul.list_id = $%d", 
	
	todoListTable, changes, usersListTable, args, args + 1)

	placeholders = append(placeholders, usId, lstId)

	_, err := r.db.Exec(query, placeholders...)


	return err
}
