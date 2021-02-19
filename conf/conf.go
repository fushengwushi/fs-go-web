package conf

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"time"
)

var BaseConf Config

type LogConf struct {
	Dir string `yaml:"dir"`
}

type RedisConf struct {
	Addr string `yaml:"addr"`
	Password string `yaml:"password"`
	PoolSize int `yaml:"poolSize"`
	DialTimeout time.Duration `yaml:"dialTimeout"`
	ReadTimeout time.Duration `yaml:"readTimeout"`
	WriteTimeout time.Duration `yaml:"writeTimeout"`
}

type Config struct {
	Log LogConf `yaml:"log"`
	Redis RedisConf `yaml:"redis"`
}

func InitConf()  {
	confPath := "./conf/config.yaml"
	if yamlFile, err := ioutil.ReadFile(confPath); err != nil {
		panic("read conf error: " + err.Error())
	} else if err = yaml.Unmarshal(yamlFile, &BaseConf); err != nil {
		panic("conf file unmarshal error: " + err.Error())
	}
}