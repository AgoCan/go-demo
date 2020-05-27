package main

import (
	"fmt"

	"github.com/spf13/viper"
)

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
	fmt.Println(viper.Get("db.name"))
}
