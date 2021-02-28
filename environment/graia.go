package environment

import (
	"os"
	"os/exec"
)

func CheckGraia() {
	_, err := exec.LookPath("graiax")
	if err != nil {
		cmd := exec.Command("pip", "install", "-U", "graiax_cli")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}
}
