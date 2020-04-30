package t

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var (
	MysqlName     string
	MysqlPassword string
	MysqlUsername string
	MysqlPort     string
	MysqlHost     string
)

// Config mysql相关配置项
type Config struct {
	Db struct {
		Mysql struct {
			Name     string `yaml:"name"`
			Password string `yaml:"password"`
			Username string `yaml:"username"`
			Port     string `yaml:"port"`
			Host     string `yaml:"host"`
		}
	}
}

var C Config

func (m *Config) getConfig() *Config {
	configYaml, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		fmt.Printf("err %v\n", err)
	}
	err = yaml.Unmarshal(configYaml, m)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return m
}

func init() {
	C.getConfig()
	MysqlHost = C.Db.Mysql.Host
}
