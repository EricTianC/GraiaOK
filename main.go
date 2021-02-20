package main

import (
	"os"

	env "github.com/EricTianC/GraiaOK/environment"
)

var javaPath string

func main() {
	//检查环境
	globalES := env.NewEnvSpace()
	globalES.CheckEnv()
	mcl := globalES.MclCommand(nil)
	mcl.Stdin = os.Stdin
	mcl.Stdout = os.Stdout
	mcl.Stderr = os.Stderr
	mcl.Run()
}
