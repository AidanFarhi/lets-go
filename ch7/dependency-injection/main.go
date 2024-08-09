package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

// logging
type Logger interface {
	Log(message string)
}

type LoggerAdapter func(message string)

func (la LoggerAdapter) Log(message string) {
	la(message)
}

func LogMessage(message string) {
	fmt.Println(message)
}

// repository
type UserRepository interface {
	GetUserNameForId(userId string) (string, bool)
}

type SimpleUserRepository struct {
	repo map[string]string
}

func (sur SimpleUserRepository) GetUserNameForId(userId string) (string, bool) {
	userName, ok := sur.repo[userId]
	if !ok {
		return "", false
	}
	return userName, true
}

func NewSimpleRepository() SimpleUserRepository {
	return SimpleUserRepository{
		repo: map[string]string{
			"1": "bob",
			"2": "jim",
			"3": "steve",
		},
	}
}

// service
type MessageService interface {
	SayHello(userId string) (string, error)
	SayGoodbye(userId string) (string, error)
}

type SimpleMessageService struct {
	ur UserRepository
}

func (sms SimpleMessageService) SayHello(userId string) (string, error) {
	userName, ok := sms.ur.GetUserNameForId(userId)
	if !ok {
		return "", errors.New("userId not found")
	}
	return "Hello, " + userName, nil
}

func (sms SimpleMessageService) SayGoodbye(userId string) (string, error) {
	userName, ok := sms.ur.GetUserNameForId(userId)
	if !ok {
		return "", errors.New("userId not found")
	}
	return "Goodbye, " + userName, nil
}

func NewSimpleMessageService(ur UserRepository) SimpleMessageService {
	return SimpleMessageService{
		ur: ur,
	}
}

// controller
type Controller struct {
	s MessageService
	l Logger
}

func (c Controller) SayHello(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")
	w.Header().Add("content-type", "application/json")
	msg, err := c.s.SayHello(userId)
	if err != nil {
		c.l.Log("error while handling userId: " + userId)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"msg": "error handling request"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"msg": "` + msg + `"}`))
}

func (c Controller) SayGoodbye(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")
	w.Header().Add("content-type", "application/json")
	msg, err := c.s.SayGoodbye(userId)
	if err != nil {
		c.l.Log("error while handling userId: " + userId)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"msg": "error handling request"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"msg": "` + msg + `"}`))
}

func NewController(s MessageService, l Logger) Controller {
	return Controller{
		s: s,
		l: l,
	}
}

func main() {
	logger := LoggerAdapter(LogMessage)
	repository := NewSimpleRepository()
	service := NewSimpleMessageService(repository)
	controller := NewController(service, logger)
	http.HandleFunc("/hello", controller.SayHello)
	http.HandleFunc("/goodbye", controller.SayGoodbye)
	log.Fatal(http.ListenAndServe(":5555", nil))
}
