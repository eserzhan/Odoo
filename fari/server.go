package todo

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	server *http.Server
}

func NewServer() *Server{
	return &Server{}
}


func (s *Server) Run (port string, handler *gin.Engine) error {
	s.server = &http.Server{
		Addr:           ":" + port,
		Handler: handler,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second}

	return s.server.ListenAndServe()
}