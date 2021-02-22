package environment

import (
	"fmt"
	"log"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"sync"

	down "github.com/EricTianC/GraiaOK/download"
)

const (
	PY_MIRROR  = "https://npm.taobao.org/mirrors/python"
	PY_VERSION = "3.9.2"
)

var pyse = &SimEnv{
	Name:     "Python",
	ExecName: "python",
	BasePath: "python",
}

func (es *EnvSpace) CheckPython(pys chan<- bool) {
	if pyse.IsInstalled() {
		ver, _ := getPyVersion()
		verNum, _ := strconv.ParseFloat(ver[:3], 32)
		if verNum < 3.8 {
			pys <- false
			return
		}
		pys <- true
		return
	}

	if pyse.LookForExecFileinSpace(es) {
		pys <- true
		return
	}

	pys <- false
}

func (es *EnvSpace) DownloadPy(gwg *sync.WaitGroup, javacomplete <-chan struct{}) {
	defer gwg.Done()
	if runtime.GOOS == "windows" {
		downloadPyWindows(es)
		<-javacomplete
	} else if runtime.GOOS == "linux" {
		<-javacomplete
		downloadPyLinux()
	} else {
		<-javacomplete
		log.Print("您的系统暂不支持，请手动配置Python3.8以上版本(含3.8)")
	}
}

func downloadPyLinux() {
	log.Println()
}

func downloadPyWindows(es *EnvSpace) {
	var downUrl string
	if runtime.GOARCH == "amd64" {
		downUrl = fmt.Sprintf("%s/%s/python-%[2]s-amd64.exe", PY_MIRROR, PY_VERSION)
	} else {
		downUrl = fmt.Sprintf("%s/%s/python-%[2]s.exe", PY_MIRROR, PY_VERSION)
	}
	name := path.Base(downUrl)
	down.DownloadFile(name, downUrl, "下载Py安装包")
	targetDir, _ := filepath.Abs(filepath.Join(es.BasePath, pyse.BasePath))
	cmd := exec.Command("./"+name, "/passive", "TargetDir="+targetDir, "PrependPath=1")
	cmd.Run()
	pyse.LookForExecFileinSpace(es)

}

func getPyVersion() (string, error) {
	cmd := exec.Command(pyse.ExecName, "--version")
	res, err := cmd.Output()
	if err != nil {
		return "", err
	}
	re, _ := regexp.Compile("\\d{1,2}\\.\\d{1,2}\\.\\d{1,2}")
	return string(re.Find(res)), nil
}
