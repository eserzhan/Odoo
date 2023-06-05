package handler

import (
	"net/http"
	"github.com/eserzhan/rest"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createList(c *gin.Context) {
	var list todo.Todo_lists

	userId, ok := c.Get("userId")

	if !ok {
		errorResponse(c, http.StatusInternalServerError, "userId not found")
		return 
	}
	
	id, ok := userId.(int)

	if !ok {
		errorResponse(c, http.StatusInternalServerError, "invalid type of userId")
		return 
	}

	if err := c.ShouldBindJSON(&list); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	listId, err := h.service.TodoList.Create(id, list)

	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return 
	}

	c.JSON(200, map[string]interface{}{
		"id": listId,
	})
}

func (h *Handler) getList(c *gin.Context){
	var list []todo.Todo_lists

	userId, ok := c.Get("userId")

	if !ok {
		errorResponse(c, http.StatusInternalServerError, "userId not found")
		return 
	}
	
	id, ok := userId.(int)

	if !ok {
		errorResponse(c, http.StatusInternalServerError, "invalid type of userId")
		return 
	}


	list, err := h.service.TodoList.Get(id)

	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return 
	}

	c.JSON(200, map[string]interface{}{
		"list": list,
	})
}

func (h *Handler) getListById(c *gin.Context){
	

	userId, ok := c.Get("userId")

	if !ok {
		errorResponse(c, http.StatusInternalServerError, "userId not found")
		return 
	}
	
	usid, ok := userId.(int)

	if !ok {
		errorResponse(c, http.StatusInternalServerError, "invalid type of userId")
		return 
	}

	lstId := c.Param("id")

	var list todo.Todo_lists
	list, err := h.service.TodoList.GetById(usid, lstId)

	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return 
	}

	c.JSON(200, map[string]interface{}{
		"list": list,
	})
}

func (h *Handler) deleteList(c *gin.Context){
	

	userId, ok := c.Get("userId")

	if !ok {
		errorResponse(c, http.StatusInternalServerError, "userId not found")
		return 
	}
	
	usid, ok := userId.(int)

	if !ok {
		errorResponse(c, http.StatusInternalServerError, "invalid type of userId")
		return 
	}

	lstId := c.Param("id")

	err := h.service.TodoList.Delete(usid, lstId)

	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return 
	}

	c.JSON(200, map[string]interface{}{
		"list": "ok",
	})
}

func (h *Handler) updateListById(c *gin.Context){
	

	userId, ok := c.Get("userId")

	if !ok {
		errorResponse(c, http.StatusInternalServerError, "userId not found")
		return 
	}
	
	usid, ok := userId.(int)

	if !ok {
		errorResponse(c, http.StatusInternalServerError, "invalid type of userId")
		return 
	}

	lstId := c.Param("id")

	var list todo.UpdateTodoLists

	if err := c.ShouldBindJSON(&list); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.service.TodoList.Update(usid, lstId, list)

	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return 
	}

	c.JSON(200, map[string]interface{}{
		"result": "ok",
	})
}