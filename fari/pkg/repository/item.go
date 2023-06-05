package repository

import (
	"fmt"
	"strconv"
	"strings"

	//"strings"

	"github.com/eserzhan/rest"
	"github.com/jmoiron/sqlx"
)

const (
	todoItemTable = "todo_items"

	listsItemTable = "lists_items"
)

type TodoItemRepository struct {
	db *sqlx.DB
}

func NewTodoItemRepository(db *sqlx.DB) *TodoItemRepository {
	return &TodoItemRepository{db: db}
}

func (r *TodoItemRepository) Create(lstId string, item todo.Todo_items) (int, error){
	
	tx, err := r.db.Begin()

	if err != nil {
		return 0, err 
	}
	defer tx.Rollback()

	var itemId int 
	query_1 := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoItemTable)
	err = tx.QueryRow(query_1, item.Title, item.Description).Scan(&itemId)

	if err != nil {
		return 0, err 
	}
	query_2 := fmt.Sprintf("INSERT INTO %s (item_id, list_id) VALUES ($1, $2)", listsItemTable)
	_, err = tx.Exec(query_2, itemId, lstId)

	if err != nil {
		return 0, err
	}

	return itemId, tx.Commit() 
}

func (r *TodoItemRepository) Get(userId int, lstId string) ([]todo.Todo_items, error) {
	id := strconv.Itoa(userId)

	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description FROM %s ti JOIN %s li ON ti.id = li.item_id 
						JOIN %s ul ON li.list_id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2`, 
						todoItemTable, listsItemTable, usersListTable)

	var res []todo.Todo_items
	err := r.db.Select(&res, query, id, lstId)

	
	if err != nil {
		return []todo.Todo_items{}, err
	}

	return res, nil
}

func (r *TodoItemRepository) GetById(userId int, itemId string) (todo.Todo_items, error) {
	usId := strconv.Itoa(userId)

	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti 
	JOIN %s li ON ti.id = li.item_id JOIN %s ul ON li.list_id = ul.list_id WHERE ti.id = $1 and ul.user_id = $2`, todoItemTable, listsItemTable, usersListTable)

	var res todo.Todo_items
	err := r.db.Get(&res, query, itemId, usId)

	if err != nil {
		return todo.Todo_items{}, err
	}

	return res, nil
}

func (r *TodoItemRepository) Delete(userId int, itemId string) (error) {
	usId := strconv.Itoa(userId)

	//query := fmt.Sprintf("DELETE FROM %s ul USING %s tl WHERE tl.id = ul.Item_id and ul.user_id = $1 and ul.Item_id = $2", todoItemTable, usersItemTable)
	query := fmt.Sprintf("DELETE FROM %s ti USING %s li, %s ul WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id=$1 AND li.item_id=$2",
	todoItemTable, listsItemTable, usersListTable)


	_, err := r.db.Exec(query, usId, itemId)


	return err
}


func (r *TodoItemRepository) Update(userId int, itemId string, up todo.UpdateTodoItems) (error) {
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

	if up.Done != nil {
		fields = append(fields, fmt.Sprintf("done=$%d", args))
		placeholders = append(placeholders, *up.Done)
		args += 1
	}
	
	changes := strings.Join(fields, ", ")

	query := fmt.Sprintf("UPDATE %s ti SET %s FROM %s li JOIN %s ul ON li.list_id = ul.list_id WHERE ti.id = li.item_id AND ul.user_id = $%d AND li.item_id = $%d", 
	
	todoItemTable, changes, listsItemTable, usersListTable, args, args + 1)

	placeholders = append(placeholders, usId, itemId)

	_, err := r.db.Exec(query, placeholders...)


	return err
}
