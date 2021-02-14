package main

import "os/exec"

const JAVA = "Java"

func checkJRE() {
	if checkJavaBin() {
		return
	}

}

func checkJavaBin() bool {
	_, err := exec.LookPath(JAVA)
	if err != nil {
		return false
	}
	return true
}
