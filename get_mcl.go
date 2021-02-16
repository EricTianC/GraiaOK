package main

import (
	"log"
	"os"
)

const (
	REPOURL = "iTXTech/mirai-console-loader"
	MCL_ZIP = "mcl.zip"
)

func get_mcl() {
	if _, err := os.Stat(MCL_ZIP); err != nil {
		download_mcl()
	}
	err := unpack(MCL_ZIP, ".")
	if err != nil {
		log.Panicf("解压失败：%s", err)
	}
	os.Remove(MCL_ZIP)
}

func download_mcl() {
	downUrl, err := get_latest_version_url(REPOURL, 0)
	if err != nil {
		log.Panic(err)
	}
	err = downloadFile(MCL_ZIP, downUrl)
	if err != nil {
		log.Panicf("下载失败：%s", err)
	}
}
