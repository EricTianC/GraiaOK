package environment

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	down "github.com/EricTianC/GraiaOK/download"
)

const (
	REPOURL = "iTXTech/mirai-console-loader"
	MCL_ZIP = "mcl.zip"
)

var mclse = &SimEnv{
	Name:     "mcl",
	BasePath: "mcl",
	ExecName: "mcl.jar",
}

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
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	cmd.Stderr = os.Stderr
	err = cmd.Start()
	if err != nil {
		return err
	}
	finished := make(chan struct{})
	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			text := scanner.Text()
			if strings.Contains(text, "mirai-console started successfully.") {
				finished <- struct{}{}
			}
		}
	}()
	go func() {
		defer stdin.Close()
		<-finished
		stdin.Write([]byte("stop\n"))
	}()
	cmd.Wait()
	mahargs := strings.Split("--update-package net.mamoe:mirai-api-http --channel stable --type plugin", " ")
	cmd = es.MclCommand(mahargs)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	return nil
}

func (es *EnvSpace) MclCommand(args []string) *exec.Cmd {
	if args == nil {
		args = []string{"-jar", mclse.ExecName}
	} else {
		args = append([]string{"-jar", mclse.ExecName}, args...)
	}
	cmd := exec.Command("java", args...)
	cmd.Env = append(cmd.Env, es.Envs()...)
	cmd.Dir = filepath.Join(es.BasePath, mclse.BasePath)
	return cmd
}
