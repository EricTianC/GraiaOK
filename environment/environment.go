//处理各种环境
package environment

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
)

var gwg sync.WaitGroup

//环境空间
type EnvSpace struct {
	BasePath string   //基目录
	EnvList  []SimEnv //环境列表
}

//单个环境
type SimEnv struct {
	Name, //名称
	ExecName, //可执行文件名称
	BasePath, //基目录，相对路径，如"java"
	ExecPath string //可执行文件所在的路径
}

func NewEnvSpace() *EnvSpace {
	es := &EnvSpace{
		BasePath: "./.env/",
	}
	if _, err := os.Stat(es.BasePath); err != nil {
		os.Mkdir(es.BasePath, os.ModePerm)
	}
	return es
}

func (es *EnvSpace) CheckEnv() {
	javaExits := make(chan bool)
	pyExits := make(chan bool)
	var js, pys bool
	go es.CheckJava(javaExits)
	go es.CheckPython(pyExits)
	gwg.Add(2)
	go func() {
		js = <-javaExits
		gwg.Done()
	}()
	go func() {
		pys = <-pyExits
		gwg.Done()
	}()
	gwg.Wait()
	close(javaExits)
	close(pyExits)
	complete := make(chan struct{})
	defer close(complete)
	if js == false && yesorNot("未检测到Java环境，是否下载", true) {
		gwg.Add(1)
		es.DownloadJava(&gwg, complete)
	} else {
		go func() {
			complete <- struct{}{}
		}()
	}
	if pys == false && yesorNot("未检测到Python环境或版本过低，是否下载", true) {
		gwg.Add(1)
		es.DownloadPy(&gwg, complete)
	} else {
		go func() {
			<-complete
		}()
	}
	gwg.Wait()
	err := es.CheckMcl()
	if err != nil {
		log.Panic(err)
	}
	CheckGraia()
}

//检查是否已安装
func (se *SimEnv) IsInstalled() bool {
	_, err := exec.LookPath(se.ExecName)
	//exec.LookPath有时候对如win10商店版Python之类的检测不到
	//这个时候就直接当没有Python环境好了反正Command也执行不了😑
	//我想我会在Readme里面写明的
	if err != nil {
		return false
	}
	return true
}

//检查在环境空间中是否有单个环境的目录
func (se *SimEnv) HasDirinEnvSpace(es *EnvSpace) bool {
	_, err := os.Stat(filepath.Join(es.BasePath, se.BasePath))
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

//在空间中查找对应可执行文件
func (se *SimEnv) LookForExecFileinSpace(es *EnvSpace) bool {
	if !se.HasDirinEnvSpace(es) {
		os.Mkdir(filepath.Join(es.BasePath, se.BasePath), os.ModePerm)
		return false
	}

	err := filepath.Walk(filepath.Join(es.BasePath, se.BasePath),
		func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() && (info.Name() == se.ExecName || info.Name() == se.ExecName+".exe") {
				se.ExecPath = filepath.Dir(path)
			}
			return err
		})
	if err != nil {
		return false
	}
	es.EnvList = append(es.EnvList, *se)
	return true
}

func (es *EnvSpace) Envs() []string {
	if len(es.EnvList) == 0 {
		return []string{""}
	}
	var envs []string
	for _, env := range es.EnvList {
		envs = append(envs, env.ExecPath)
	}
	return envs
}

func yesorNot(question string, defau bool) bool {
	if defau == false {
		fmt.Print(question + "[y/n](默认n)：")
	} else {
		fmt.Print(question + "[y/n]：")
	}
	var opt string
	for {
		fmt.Scanln(&opt)
		switch opt {
		case "y", "Y":
			return true

		case "n", "N":
			return false
		case "":
			return defau
		default:
			fmt.Println("输入不正确")
		}
	}
}
