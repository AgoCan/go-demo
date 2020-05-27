package main

import (
	"fmt"

	"github.com/agocan/go-demo/2020-05/gin-zap-demo/zap"
	"github.com/gin-gonic/gin"
)

func main() {
	c := zap.LogConfig{}
	c.Filename = "ttt.log"
	c.Level = "info"
	c.MaxSize = 10
	c.MaxAge = 10
	c.MaxBackups = 5
	if err := zap.InitLogger(&c); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	r := gin.New()
	r.Use(zap.GinLogger(zap.Logger), zap.GinRecovery(zap.Logger, true))
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "hellp world")
	})
	r.Run()
}
