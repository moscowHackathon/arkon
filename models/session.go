package models

import (
	"errors"
	"sync"
	"github.com/astaxie/beego"
)

type Session struct {
	ID string
	calc *SessionCalc
}

func (s Session) GetQuestion() string {
	if ok,i := s.calc.GetNextToCheck(); ok {
		return features[i]
	}
	return "Error: no questions"
}

func (s Session) Answer(a int) (bool, string) {

	beego.Error("session answer: ", a)
	s.calc.ApplyAnswer(getValue(a))
	res,i := s.calc.CheckStatus()

	if res {
		if i>=0 {
			return true, teams[i]
		}
		return true, ""
	}

	return false, ""
}

func getValue(a int)int {
	switch a{
	case 1:
		return YES
	case -1:
		return NO
	default:
		return NA
	}
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
