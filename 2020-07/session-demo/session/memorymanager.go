package session

import (
	"sync"

	uuid "github.com/satori/go.uuid"
)

// MemoryManager 内存
type MemoryManager struct {
	sessionMap map[string]Session
	rwlock     sync.RWMutex
}

// NewMemoryManager new一个实例
func NewMemoryManager() Manager {
	sr := &MemoryManager{
		sessionMap: make(map[string]Session, 1024),
	}

	return sr
}

// Init 初始化实例
func (s *MemoryManager) Init(addr string, options ...string) (err error) {
	return
}

// Get 获取session
func (s *MemoryManager) Get(sessionID string) (session Session, err error) {
	s.rwlock.RLock()
	defer s.rwlock.RUnlock()

	session, ok := s.sessionMap[sessionID]
	if !ok {
		err = ErrSessionNotExist
		return
	}

	return
}

// CreateSession 创建session
func (s *MemoryManager) CreateSession() (session Session, err error) {
	s.rwlock.Lock()
	defer s.rwlock.Unlock()
	id := uuid.NewV4()
	if err != nil {
		return
	}

	sessionID := id.String()
	session = NewMemory(sessionID)

	s.sessionMap[sessionID] = session
	return
}
