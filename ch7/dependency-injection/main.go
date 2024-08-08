package main

import (
	"errors"
	"fmt"
	"net/http"
)

// interfaces
type DataStore interface {
	UserNameForID(userId string) (string, bool)
}

type Logger interface {
	Log(message string)
}

type MessageService interface {
	SayHello(userId string) (string, error)
	SayGoodbye(userId string) (string, error)
}

// adapters
type LoggerAdapter func(message string)

func (la LoggerAdapter) Log(message string) {
	la(message)
}

// implementations
func LogOutput(message string) {
	fmt.Println(message)
}

type SimpleDataStore struct {
	userData map[string]string
}

func (sds SimpleDataStore) UserNameForID(userId string) (string, bool) {
	name, ok := sds.userData[userId]
	return name, ok
}

func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		userData: map[string]string{
			"1": "bob",
			"2": "joe",
			"3": "jim",
		},
	}
}

type SimpleMessageService struct {
	l  Logger
	ds DataStore
}

func (ms SimpleMessageService) SayHello(userId string) (string, error) {
	ms.l.Log("in SayHello for " + userId)
	name, ok := ms.ds.UserNameForID(userId)
	if !ok {
		return "", errors.New("uknown user")
	}
	return "Hello, " + name, nil
}

func (ms SimpleMessageService) SayGoodbye(userId string) (string, error) {
	ms.l.Log("in SayHello for " + userId)
	name, ok := ms.ds.UserNameForID(userId)
	if !ok {
		return "", errors.New("uknown user")
	}
	return "Goodbye, " + name, nil
}

func NewSimpleMessageService(l Logger, ds DataStore) SimpleMessageService {
	return SimpleMessageService{
		l:  l,
		ds: ds,
	}
}

type Controller struct {
	l  Logger
	ms MessageService
}

func (c Controller) SayHello(w http.ResponseWriter, r *http.Request) {
	c.l.Log("In Controller.SayHello")
	userId := r.URL.Query().Get("user_id")
	message, err := c.ms.SayHello(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(message))
}

func (c Controller) SayGoodbye(w http.ResponseWriter, r *http.Request) {
	c.l.Log("In Controller.SayGoodbye")
	userId := r.URL.Query().Get("user_id")
	message, err := c.ms.SayGoodbye(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(message))
}

func NewController(l Logger, ms MessageService) Controller {
	return Controller{
		l:  l,
		ms: ms,
	}
}

func main() {
	logger := LoggerAdapter(LogOutput)
	dataStore := NewSimpleDataStore()
	messageService := NewSimpleMessageService(logger, dataStore)
	controller := NewController(logger, messageService)
	http.HandleFunc("/hello", controller.SayHello)
	http.HandleFunc("/goodbye", controller.SayGoodbye)
	http.ListenAndServe(":8888", nil)
}
