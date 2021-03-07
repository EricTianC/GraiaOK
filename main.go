package main

import (
	env "github.com/EricTianC/GraiaOK/environment"
)

var javaPath string

func main() {
	//检查环境
	globalES := env.NewEnvSpace()
	globalES.CheckEnv()

	//go func() {
	mcl := globalES.MclCommand(nil)
	err := mcl.Run()
	if err != nil {
		panic(err)
	}
	//}()

	//	go func() {
	//
	//	}
}
