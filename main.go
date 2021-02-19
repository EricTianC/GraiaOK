package main

import (
	"log"

	env "github.com/EricTianC/GraiaOK/environment"
)

var javaPath string

func main() {
	//检查环境
	globalES := env.NewEnvSpace()
	globalES.CheckEnv()
	log.Println("这边帮您配置好了Java和mcl，还请您花费宝贵的时间自己安装一下Python呢")
}
