package environment

import (
	"path/filepath"
	"runtime"
)

const (
	MIRROR       = "https://mirrors.tuna.tsinghua.edu.cn/AdoptOpenJDK/%s/%s/%s/%s/" //镜像地址，目前使用清华源，如要更换源需修改格式化代码
	JAVA_VERSION = "15"                                                             //要下载的Java版本
	JDK_OR_JRE   = "jre"                                                            //下载JDK还是JRE，然而清华原貌似没的选（Jre也有javac）
)

//因为镜像源的目录命名和runtime.GOOS不同，所以需要映射一下
var OS = map[string]string{
	"windows": "windows",
	"linux":   "linux",
	"macos":   "mac",
}

//同上
var ARCH = map[string]string{
	"amd64": "x64",
	"386":   "x32",
	"arm64": "arm",
	"arm":   "arm",
}

var jenv = &SimEnv{
	Name:     "Java",
	ExecName: "java",
	BasePath: "java",
}

func (es *EnvSpace) CheckJavaEnv() {

	if jenv.IsInstalled() { //之所以把IsInstalled单独列成函数，是因为有的时候需要优先本地已装好的环境，有的时候又要优先自动装的环境
		return
	}

	//检查EnvSpace中是否有Java可执行文件
	//err := filepath.Walk(filepath.Join(es.BasePath, jenv.BasePath), findJava)
	switch runtime.GOOS {
	case "macos":
		matches, err := filepath.Glob(filepath.Join(es.BasePath, jenv.BasePath, "*/Contents/Home/bin/*"))
		if err == nil && len(matches) != 0 {
			jenv.ExecPath = filepath.Dir(matches[0])
		}
	default:
		matches, err := filepath.Glob("./jre/*/bin/*")
		if err == nil && len(matches) != 0 {
			jenv.ExecPath = filepath.Dir(matches[0])
		}
	}

}

//寻找java可执行文件
/*
func findJava(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		filepath.Walk(filepath.Join(path, info.Name()), findJava)
	} else if info.Name() == "java" {
		jenv.ExecPath, _ = filepath.Abs(path)
		return nil
	}
	return err
}
*/
