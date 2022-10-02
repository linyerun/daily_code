package redis

import (
	"gopkg.in/ini.v1"
	"strconv"
)

var (
	Addr string
	Pwd  string
	Port int
)

func InitRedisConfig() {
	file, err := ini.Load("static/config.ini")
	if err != nil {
		panic(err)
	}
	Addr = file.Section("redis").Key("addr").String()
	Pwd = file.Section("redis").Key("password").String()
	Port, err = strconv.Atoi(file.Section("redis").Key("port").String())
	if err != nil {
		panic("redis的port输入有误，" + err.Error())
	}
}
