package environment

import (
	"log"
	"runtime"
	"sync"
)

var pyse = &SimEnv{
	Name:     "Python",
	ExecName: "python",
	BasePath: "python",
}

func (es *EnvSpace) CheckPython(pys chan<- bool) {
	if pyse.IsInstalled() {
		pys <- true
		return
	}

	if pyse.LookForExecFileInSpace(es) {
		pys <- true
		return
	}

	pys <- false
}

func (es *EnvSpace) DownloadPy(gwg *sync.WaitGroup, javacomplete <-chan struct{}) {
	defer gwg.Done()
	if runtime.GOOS == "windows" {
		downloadPyWindows()
		<-javacomplete
	} else if runtime.GOOS == "linux" {
		<-javacomplete
		downloadPyLinux()
	} else {
		<-javacomplete
		log.Print("您的系统暂不支持，请手动配置Python3.8以上版本")
	}
}

func downloadPyLinux() {
	log.Println()
}

func downloadPyWindows() {

}
