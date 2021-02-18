package main

import (
	"fmt"

	env "github.com/EricTianC/GraiaOK/environment"
)

var javaPath string

/*
func main() {
	checkJRE()
	check_mcl()
	args := []string{"-jar", "mcl.jar"}
	args = append(args, os.Args...)
	cmd := exec.Command(filepath.Join(javaPath, JAVA), args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}
*/

func main() {
	globalES := env.NewEnvSpace()
	globalES.CheckEnv()
	fmt.Print(globalES.EnvList)
}
