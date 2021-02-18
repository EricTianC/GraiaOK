//处理各种环境
package environment

import (
	"os/exec"
)

//环境空间
type EnvSpace struct {
	BasePath string   //基目录
	EnvList  []SimEnv //环境列表
}

//单个环境
type SimEnv struct {
	Name     string //名称
	ExecName string //可执行文件名称
	BasePath string //基目录，相对路径，如"java"
	ExecPath string //可执行文件所在的路径
}

func NewEnvSpace() *EnvSpace {
	return &EnvSpace{
		BasePath: ".",
	}
}

func (es *EnvSpace) CheckEnv() {
	es.CheckJavaEnv()
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
