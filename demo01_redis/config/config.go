package config

import (
	"gopkg.in/ini.v1"
	"strconv"
)

var (
	Addr      string
	Pwd       string
	Port      int
	EmailFrom string
	EmailPwd  string
)

// initConfig 初始化配置
func initConfig() {
	file, err := ini.Load("public/config.ini")
	if err != nil {
		panic(err)
	}
	Addr = file.Section("redis").Key("addr").String()
	Pwd = file.Section("redis").Key("password").String()
	p := file.Section("redis").Key("port").String()
	if port, err := strconv.Atoi(p); err == nil {
		Port = port
	} else {
		panic(err)
	}
	EmailFrom = file.Section("email").Key("from").String()
	EmailPwd = file.Section("email").Key("pwd").String()
}
