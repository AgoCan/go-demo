package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type config struct {
	Db struct {
		Name     string
		Password string
	}
}

type config2 struct {
	Log struct {
		Path     string
		Password string
	}
}

var c2 *config2
var c *config

func main() {
	// viper.SetConfigName("config") // 文件名
	// viper.SetConfigType("yaml")  // 类型
	// viper.AddConfigPath(".")  // 扫描目录。不在一文件后缀

	// 一步到位，直接读取文件
	viper.SetConfigFile("./config.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("err:", err)
	}
	err = viper.Unmarshal(&c)
	if err != nil {
		fmt.Println(err)
	}
	err = viper.Unmarshal(&c2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c, c2)
}
