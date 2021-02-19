package main

import (
	env "github.com/EricTianC/GraiaOK/environment"
)

var javaPath string

func main() {
	//检查环境
	globalES := env.NewEnvSpace()
	globalES.CheckEnv()
}
