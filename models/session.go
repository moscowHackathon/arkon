package models

import (
	"errors"
	"sync"
)

type Session struct {
	ID string
}

func (s Session) GetQuestion() string {
	return "Question 1"
}

type SessionMap map[string]Session

var sessions = make(SessionMap)
var sessionMutex sync.RWMutex

func NewSession(id string) Session {
	s := Session{ID: id}

	sessionMutex.Lock()
	sessions[id] = s
	sessionMutex.Unlock()

	return s
}

func GetSession(id string) (Session, error) {
	s, b := sessions[id]

	if b == false {
		return s, errors.New("Session with id '" + id + "' not found")
	}

	return s, nil
}
