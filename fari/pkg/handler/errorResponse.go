package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func errorResponse(c *gin.Context, code int, err string) {
	log.Println(err)
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
}