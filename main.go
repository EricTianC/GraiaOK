package main

import (
	"os"
	"os/exec"
)

var javaPath string

func main() {
	checkJRE()
	check_mcl()
	args := []string{"-jar", "mcl.jar"}
	args = append(args, os.Args...)
	cmd := exec.Command(javaPath+JAVA, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}
