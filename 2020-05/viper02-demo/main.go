package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var config *viper.Viper

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init(env string) {
	var err error
	config = viper.New()
	config.AutomaticEnv()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("../config/")
	config.AddConfigPath("config/")
	err = config.ReadInConfig()
	if err != nil{fmt.Println(err)}
}

func relativePath(basedir string, path *string) {
	p := *path
	if len(p) > 0 && p[0] != '/' {
		*path = filepath.Join(basedir, p)
	}
}

func GetConfig() *viper.Viper {
	return config
}
func test01(){
	c := GetConfig()
	a := c.GetString("TTT")
	fmt.Println(a)

	os.Setenv("TTT", "abc")
	e := c.Get("TTT")
	fmt.Println(e)
	//读取已经加载到default中的环境变量
	if env := viper.Get("JAVA_HOME"); env == nil {
		println("error!")
	} else {
		fmt.Printf("%#v\n", env)
	}
}

func test02(){
	// 变量不能有小数点，所以暂时娶不到
	c := GetConfig()
	a := c.GetString("eC")
	fmt.Println(a)
}

func main(){
	Init("e")
	test02()
}
