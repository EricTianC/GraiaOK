package environment

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	down "github.com/EricTianC/GraiaOK/download"
)

const (
	REPOURL = "iTXTech/mirai-console-loader"
	MCL_ZIP = "mcl.zip"
)

var (
	mclse = &SimEnv{
		Name:     "mcl",
		BasePath: "mcl",
		ExecName: "mcl.jar",
	}
	once sync.Once
)

func (es *EnvSpace) CheckMcl() error {
	if _, err := os.Stat(filepath.Join(es.BasePath, mclse.BasePath, mclse.ExecName)); err == nil {
		return nil
	}

	downUrl, err := down.GetLatestVersionUrl(REPOURL, 0)
	if err != nil {
		return fmt.Errorf("无法获取版本信息：%v", err)
	}

	err = down.DownloadFile(MCL_ZIP, downUrl, "下载mcl")
	if err != nil {
		return fmt.Errorf("下载失败：%v", err)
	}

	err = down.Unpack(MCL_ZIP, filepath.Join(es.BasePath, mclse.BasePath))
	if err != nil {
		return fmt.Errorf("解压失败：%v", err)
	}

	err = firstRunMcl(es)
	if err != nil {
		return fmt.Errorf("第一次运行mcl错误：%v", err)
	}
	return nil
}

func firstRunMcl(es *EnvSpace) error {
	cmd := es.MclCommand(nil)

	//因mcl原因无法获取
	//stdin, err := cmd.StdinPipe()
	//if err != nil {
	//	return fmt.Errorf("获取Stdin管道错误：%v", err)
	//}
	log.Println("请在mcl自动下载完毕后手动输入stop并回车，后面会自动配置mah")
	cmd.Run()
	mahargs := strings.Split("--update-package net.mamoe:mirai-api-http --channel stable --type plugin", " ")
	cmd = es.MclCommand(mahargs)
	cmd.Run()
	return nil
}

//Stdin已强制设定
func (es *EnvSpace) MclCommand(args []string) *exec.Cmd {
	if args == nil {
		args = []string{"-jar", mclse.ExecName}
	} else {
		args = append([]string{"-jar", mclse.ExecName}, args...)
	}
	javaSimEnv, err := es.FindSimEnv("Java")
	var javaPath string
	if err != nil {
		javaPath, _ = exec.LookPath("Java")
	} else {
		javaPath = javaSimEnv.ExecPath
	}
	javaPath, _ = filepath.Abs(javaPath)
	cmd := exec.Command(filepath.Join(javaPath, "java"), args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// once.Do(func() {
	// 	var sep string //分隔符
	// 	switch runtime.GOOS {
	// 	case "windows":
	// 		sep = ";"
	// 	default:
	// 		sep = ":"
	// 	}
	// 	os.Setenv("PATH", os.Getenv("PATH")+sep+strings.Join(es.Envs(), sep))
	// })
	cmd.Dir = filepath.Join(es.BasePath, mclse.BasePath)
	return cmd
}
