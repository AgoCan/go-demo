package session

import "errors"

var (
	// ErrSessionNotExist session 不存在
	ErrSessionNotExist = errors.New("session not exists")
	// ErrKeyNotExistInSession session的key不存在
	ErrKeyNotExistInSession = errors.New("key not exists in session")
)
