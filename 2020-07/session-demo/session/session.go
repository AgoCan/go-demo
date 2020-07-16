package session

// Session 接口
type Session interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
	Del(key string) error
	Save() error
}

// Manager session管理
type Manager interface {
	Init(addr string, options ...string) (err error)
	CreateSession() (session Session, err error)
	Get(sessionID string) (session Session, err error)
}
