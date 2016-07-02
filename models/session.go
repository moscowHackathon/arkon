package models

import (
	"errors"
	"sync"
)

type Session struct {
	ID string
	calc SessionCalc
}

func (s Session) GetQuestion() string {
	i := s.calc.GetNextToCheck()
	return features[i]
}

func (s Session) Answer(a int) bool {
	return false
}

type SessionMap map[string]Session

var sessions = make(SessionMap)
var sessionMutex sync.RWMutex

func NewSession(id string) Session {
	s := Session{ID: id, calc: newSessionCalc(cases)}

	sessionMutex.Lock()
	sessions[id] = s
	sessionMutex.Unlock()

	return s
}

func GetSession(id string) (Session, error) {

	sessionMutex.RLock()
	s, b := sessions[id]
	sessionMutex.RUnlock()

	if b == false {
		return s, errors.New("Session with id '" + id + "' not found")
	}

	return s, nil
}
