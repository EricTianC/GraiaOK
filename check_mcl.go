package main

/*
import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	REPOURL = "iTXTech/mirai-console-loader"
	MCL_ZIP = "mcl.zip"
)

var MCL_BASE_ARG = []string{"-jar", "mcl.jar"}

func check_mcl() {
	if _, err := os.Stat("mcl.jar"); err != nil {
		get_mcl()
		first_run_mcl()
	}
}

func get_mcl() {
	if _, err := os.Stat(MCL_ZIP); err != nil {
		download_mcl()
	}
	err := unpack(MCL_ZIP, ".")
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

func first_run_mcl() {
	args := append(MCL_BASE_ARG, os.Args...)
	cmd := exec.Command(filepath.Join(javaPath, JAVA), args...)
	cmd.Stdin = os.Stdin
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Stderr = os.Stderr
	err = cmd.Start()
	if err != nil {
		log.Panic(err)
	}
	go wait_first_complete(stdout)
	cmd.Wait()
	args = append(MCL_BASE_ARG, "--update-package net.mamoe:mirai-api-http --channel stable --type plugin")
	cmd = exec.Command(filepath.Join(javaPath, JAVA), args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func wait_first_complete(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, "mirai-console started successfully.") {
			log.Println("麻烦您手动输一下stop并回车，谢谢了")
		}
		fmt.Println(text)
	}
}
*/
