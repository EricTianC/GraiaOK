package environment

import (
	"fmt"
	"os"
	"path/filepath"

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

	err = down.DownloadFile(MCL_ZIP, downUrl)
	if err != nil {
		return fmt.Errorf("下载失败：%v", err)
	}

	err = down.Unpack(MCL_ZIP, filepath.Join(es.BasePath, mclse.BasePath))
	if err != nil {
		return fmt.Errorf("解压失败：%v", err)
	}

	err = firstRunMcl()
	if err != nil {
		return fmt.Errorf("第一次运行mcl错误：%v", err)
	}
	return nil
}

func firstRunMcl() error {

	return nil
}
