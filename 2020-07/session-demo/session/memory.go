package session

import (
	"sync"
)

// Memory 内存的session
type Memory struct {
	data   map[string]interface{}
	id     string
	rwlock sync.RWMutex
}

// NewMemory 实例出一个新的session
func NewMemory(id string) *Memory {
	s := &Memory{
		id:   id,
		data: make(map[string]interface{}, 8),
	}

	return s
}

// Set 设置内存session
func (m *Memory) Set(key string, value interface{}) (err error) {

	m.rwlock.Lock()
	defer m.rwlock.Unlock()

	m.data[key] = value
	return
}

// Get 根据key获取session
func (m *Memory) Get(key string) (value interface{}, err error) {
	m.rwlock.RLock()
	defer m.rwlock.RUnlock()

	value, ok := m.data[key]
	if !ok {
		err = ErrKeyNotExistInSession
		return
	}

	return
}

// Del 删除session
func (m *Memory) Del(key string) (err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()

	delete(m.data, key)
	return
}

// Save 保存
func (m *Memory) Save() (err error) {
	return
}
