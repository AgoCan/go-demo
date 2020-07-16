package session

import "fmt"

var (
	manager Manager
)

//Init 初始化session，可以使用多种方式
//1. memory， 返回一个内存的session管理类
//2. redis, 返回一个redis的session管理类
func Init(provider string, addr string, options ...string) (err error) {

	switch provider {
	case "memory":
		manager = NewMemoryManager()
	// case "redis":
	// 	manager = NewRedisSessionMgr()
	default:
		err = fmt.Errorf("not support")
		return
	}

	err = manager.Init(addr, options...)
	return
}
