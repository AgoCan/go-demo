package session

import (
	"sync"
	"time"

	"github.com/garyburd/redigo/redis"
	uuid "github.com/satori/go.uuid"
)

// RedisManager r
type RedisManager struct {
	addr       string
	passwd     string
	sessionMap map[string]Session
	pool       *redis.Pool
	rwlock     sync.RWMutex
}

// NewRedisManager 实例
func NewRedisManager() Manager {
	sr := &RedisManager{
		sessionMap: make(map[string]Session, 1024),
	}

	return sr
}

//初始化一个pool
func newPool(server, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     64,
		MaxActive:   1000,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			/*
			   if _, err := c.Do("AUTH", password); err != nil {
			   c.Close()
			   return nil, err
			   }*/
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

// Init 初始化
func (r *RedisManager) Init(addr string, options ...string) (err error) {

	if len(options) > 0 {
		r.passwd = options[0]
	}

	r.pool = newPool(addr, r.passwd)
	r.addr = addr
	return
}

// CreateSession 创建
func (r *RedisManager) CreateSession() (session Session, err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()

	id := uuid.NewV4()
	sessionID := id.String()
	session = NewRedisSession(sessionID, r.pool)

	r.sessionMap[sessionID] = session
	return
}

// Get 获取
func (r *RedisManager) Get(sessionID string) (session Session, err error) {

	r.rwlock.RLock()
	defer r.rwlock.RUnlock()

	session, ok := r.sessionMap[sessionID]
	if !ok {
		err = ErrSessionNotExist
		return
	}
	return
}
