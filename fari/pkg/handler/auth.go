package handler

import (
	"net/http"

	"github.com/eserzhan/rest"
	"github.com/gin-gonic/gin"
)


func (h *Handler) sign_up(c *gin.Context) {
	var user todo.User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}

	id, err := h.service.Authorization.CreateUser(user)

	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		//c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
	}

	c.JSON(200, gin.H{"id": id})
}

type loginUser struct {
	Password string `json:"password" binding:"required"`
    Username string `json:"username" binding:"required"`
}

func (h *Handler) sign_in(c *gin.Context) {
	var user loginUser

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}

	token, err := h.service.Authorization.GenerateToken(user.Username, user.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
	}

	c.JSON(200, gin.H{"token": token})
}

