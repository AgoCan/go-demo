package main

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

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
			_, err = c.Do("PING")
			if err != nil {
				a, err := c.Do("AUTH", password)
				if err != nil {

					fmt.Println("redis auth failed: ", err)
					return nil, err
				}
				fmt.Println(a)
			}
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

func main() {
	host := "localhost" //控制台显示的地址
	port := 6379
	server := fmt.Sprintf("%s:%d", host, port)

	r := newPool(server, "")
	c, err := r.Dial()
	if err != nil {
		return
	}
	defer c.Close()

	b, err := c.Do("SET", "keyss", "jcloud-redisss")
	if err != nil {
		fmt.Println("redis set failed: ", err)
		return
	}
	fmt.Println(b)
	l, err := c.Do("GET", "keyss")
	if err != nil {
		fmt.Println("redis get failed: ", err)
		return
	}
	fmt.Printf("--------%#v, %T\n", l, l)
	ee, err := redis.String(c.Do("GET", "keyss"))
	if err != nil {
		fmt.Println("redis get failed: ", err)
		return
	}
	fmt.Println(ee)
	//do other command...
}
