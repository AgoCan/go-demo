package session

import (
	"encoding/json"
	"sync"

	"github.com/garyburd/redigo/redis"
)

// 不知道
const (
	SessionFlagNone = iota
	SessionFlagModify
	SessionFlagLoad
)

// RedisSession sss
type RedisSession struct {
	sessionID  string
	pool       *redis.Pool
	sessionMap map[string]interface{}
	rwlock     sync.RWMutex
	flag       int
}

// NewRedisSession ss
func NewRedisSession(id string, pool *redis.Pool) *RedisSession {
	s := &RedisSession{
		sessionID:  id,
		sessionMap: make(map[string]interface{}, 8),
		flag:       SessionFlagNone,
		pool:       pool,
	}

	return s
}

// Set 设置
func (r *RedisSession) Set(key string, value interface{}) error {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()

	r.sessionMap[key] = value
	r.flag = SessionFlagModify
	return nil
}

func (r *RedisSession) loadFromRedis() (err error) {

	conn := r.pool.Get()
	reply, err := conn.Do("GET", r.sessionID)
	if err != nil {
		return
	}

	data, err := redis.String(reply, err)
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(data), &r.sessionMap)
	if err != nil {
		return
	}

	return
}

// Get 获取
func (r *RedisSession) Get(key string) (result interface{}, err error) {

	r.rwlock.RLock()
	defer r.rwlock.RUnlock()

	//实现了一个延迟加载的功能
	if r.flag == SessionFlagNone {
		//该session还没有加载，那么就从redis中加载数据
		err = r.loadFromRedis()
		if err != nil {
			return
		}
	}

	result, ok := r.sessionMap[key]
	if !ok {
		err = ErrKeyNotExistInSession
		return
	}

	return
}

// Del 删除
func (r *RedisSession) Del(key string) error {

	r.rwlock.Lock()
	defer r.rwlock.Unlock()

	r.flag = SessionFlagModify
	delete(r.sessionMap, key)
	return nil
}

// Save 保存
func (r *RedisSession) Save() (err error) {

	r.rwlock.Lock()
	defer r.rwlock.Unlock()

	if r.flag != SessionFlagModify {
		return
	}

	data, err := json.Marshal(r.sessionMap)
	if err != nil {
		return
	}

	conn := r.pool.Get()
	_, err = conn.Do("SET", r.sessionID, string(data))
	if err != nil {
		return
	}

	return
}
