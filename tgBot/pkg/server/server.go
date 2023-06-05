package server

import (
	//"context"


	"log"
	"net/http"
	"strconv"

	"github.com/eserzhan/tgBott/pkg/repository"

	"github.com/zhashkevych/go-pocket-sdk"
)

type AuthServer struct {
	server *http.Server
	client *pocket.Client
	db repository.TokenRepository
	redirectUrl string
}

func NewServer(client *pocket.Client, db repository.TokenRepository, redirectUrl string) *AuthServer{
	return &AuthServer{client: client, db: db, redirectUrl: redirectUrl}
}

func(s *AuthServer) Start() error{
	s.server = &http.Server{
		Addr: ":80",
		Handler: s,

	}

	return s.server.ListenAndServe()
}

func(s *AuthServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	chatId_str := r.URL.Query().Get("chat_id")

	if chatId_str == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	chat_id, err := strconv.ParseInt(chatId_str, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	requestToken, err := s.db.Get(chat_id, repository.RequestToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	accessToken, err := s.client.Authorize(r.Context(), requestToken)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	err = s.db.Save(chat_id, accessToken.AccessToken, repository.AccessToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	w.Header().Add("Location", s.redirectUrl)
	w.WriteHeader(http.StatusMovedPermanently)
	
	
}