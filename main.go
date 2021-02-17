package main

import (
	"fmt"
	"os"
	"os/exec"
)

var javaPath string

func main() {
	checkJRE()
	if _, err := os.Stat("mcl.jar"); err != nil {
		get_mcl()
	}
	fmt.Println(javaPath + JAVA)
	cmd := exec.Command(javaPath+JAVA, "-jar", "mcl.jar")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}
