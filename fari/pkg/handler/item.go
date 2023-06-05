package handler

import (

	"net/http"

	"github.com/eserzhan/rest"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createItem(c *gin.Context) {
	var item todo.Todo_items


	lstId := c.Param("id")

	
	
	if err := c.ShouldBindJSON(&item); err != nil {

		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	
	ItemId, err := h.service.TodoItem.Create(lstId, item)

	if err != nil {
		
		errorResponse(c, http.StatusBadRequest, err.Error())
		return 
	}

	c.JSON(200, map[string]interface{}{
		"id": ItemId,
	})
}

func (h *Handler) getItems(c *gin.Context){

	lstId := c.Param("id")

	userId, ok := c.Get("userId")

	if !ok {
		errorResponse(c, http.StatusInternalServerError, "userId not found")
		return 
	}
	
	usrId, ok := userId.(int)

	if !ok {
		errorResponse(c, http.StatusInternalServerError, "invalid type of userId")
		return 
	}

	items, err := h.service.TodoItem.Get(usrId, lstId)

	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return 
	}

	c.JSON(200, map[string]interface{}{
		"Item": items,
	})
}

func (h *Handler) getItemById(c *gin.Context){
	

	itemId := c.Param("id")

	userId, ok := c.Get("userId")

	if !ok {
		errorResponse(c, http.StatusInternalServerError, "userId not found")
		return 
	}
	
	usrId, ok := userId.(int)

	if !ok {
		errorResponse(c, http.StatusInternalServerError, "invalid type of userId")
		return 
	}

	//var item todo.Todo_items
	item, err := h.service.TodoItem.GetById(usrId, itemId)

	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return 
	}

	c.JSON(200, map[string]interface{}{
		"Item": item,
	})
}

func (h *Handler) deleteItem(c *gin.Context){
	

	itemId := c.Param("id")

	userId, ok := c.Get("userId")

	if !ok {
		errorResponse(c, http.StatusInternalServerError, "userId not found")
		return 
	}
	
	usrId, ok := userId.(int)

	if !ok {
		errorResponse(c, http.StatusInternalServerError, "invalid type of userId")
		return 
	}

	
	err := h.service.TodoItem.Delete(usrId, itemId)

	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return 
	}

	c.JSON(200, map[string]interface{}{
		"result": "ok",
	})
}

func (h *Handler) updateItem(c *gin.Context){
	

	itemId := c.Param("id")

	userId, ok := c.Get("userId")

	if !ok {
		errorResponse(c, http.StatusInternalServerError, "userId not found")
		return 
	}
	
	usrId, ok := userId.(int)

	if !ok {
		errorResponse(c, http.StatusInternalServerError, "invalid type of userId")
		return 
	}

	var item todo.UpdateTodoItems

	if err := c.ShouldBindJSON(&item); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.service.TodoItem.Update(usrId, itemId, item)

	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return 
	}

	c.JSON(200, map[string]interface{}{
		"result": "ok",
	})
}