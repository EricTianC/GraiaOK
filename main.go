package main

import (
	"os"
	"os/exec"
)

var javaPath string

func main() {
	checkJRE()
	get_mcl()
	cmd := exec.Command(javaPath+JAVA, "-jar", "mcl.jar")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}
