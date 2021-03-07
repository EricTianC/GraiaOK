package environment

import (
	"testing"
)

func TestCheckJava(t *testing.T) {
	es := NewEnvSpace()
	javaExists := make(chan bool)
	es.CheckJava(javaExists)
	_, err := es.FindSimEnv("Java")
	if err != nil {
		t.Error("并没有JavaSimEnv")
	}
}
