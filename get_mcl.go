package main

import (
	"log"
	"os"
	archiver "github.com/mholt/archiver"
)

const (
	REPOURL = "iTXTech/mirai-console-loader"
	MCL_ZIP = "mcl.zip"
)

func get_mcl() {
	if _, err := os.Stat(MCL_ZIP); err != nil {
		download_mcl()
	}
	err := archiver.Unarchive(MCL_ZIP, ".")
	if err != nil {
		log.Panicf("解压失败：%s", err)
	}
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
