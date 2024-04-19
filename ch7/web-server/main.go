package main

import (
	"errors"
	"fmt"
	"net/http"
)

type DataStore interface {
	UserNameForID(userID string) (string, bool)
}

type SimpleDataStore struct {
	userData map[string]string
}

func (sds SimpleDataStore) UserNameForID(userID string) (string, bool) {
	name, ok := sds.userData[userID]
	return name, ok
}

func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		userData: map[string]string{
			"1": "<div style=>Fred</div>",
			"2": "Joe",
			"3": "Bill",
		},
	}
}

type Logger interface {
	Log(message string)
}

type LoggerAdapter func(message string)

func (la LoggerAdapter) Log(message string) {
	la(message)
}

func LogOutput(message string) {
	fmt.Println(message)
}

type Logic interface {
	SayHello(userID string) (string, error)
}

type SimpleLogic struct {
	l  Logger
	ds DataStore
}

func (sl SimpleLogic) SayHello(userID string) (string, error) {
	sl.l.Log("in SayHello for " + userID)
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("unkown user")
	}
	return "Hello, " + name, nil
}

func (sl SimpleLogic) SayGoodbye(userID string) (string, error) {
	sl.l.Log("in SayGoodbye for" + userID)
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("unkown user")
	}
	return "Goodbye, " + name, nil
}

func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
	return SimpleLogic{
		l:  l,
		ds: ds,
	}
}

type Controller struct {
	l     Logger
	logic Logic
}

func (c Controller) SayHello(w http.ResponseWriter, r *http.Request) {
	c.l.Log("In SayHello")
	userID := r.URL.Query().Get("user_id")
	message, err := c.logic.SayHello(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(message))
}

func NewController(l Logger, logic Logic) Controller {
	return Controller{
		l:     l,
		logic: logic,
	}
}

func main() {
	logger := LoggerAdapter(LogOutput)
	dataStore := NewSimpleDataStore()
	logic := NewSimpleLogic(logger, dataStore)
	controller := NewController(logger, logic)
	http.HandleFunc("/hello", controller.SayHello)
	http.ListenAndServe(":1234", nil)
}
