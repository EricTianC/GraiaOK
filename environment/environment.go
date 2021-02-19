//处理各种环境
package environment

import (
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
	gwg.Add(1)
	go es.CheckJavaEnv(&gwg)
	gwg.Wait()
	es.CheckMcl()
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
func (se *SimEnv) LookForExecFileInSpace(es *EnvSpace) bool {
	if !se.HasDirinEnvSpace(es) {
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
