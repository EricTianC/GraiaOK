package main

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/PuerkitoBio/goquery"
)

const (
	JAVA         = "Java"
	MIRROR       = "https://mirrors.tuna.tsinghua.edu.cn/AdoptOpenJDK/%s/%s/%s/%s/"
	JAVA_VERSION = "15"
	JDK_OR_JRE   = "jre"
)

var OS = map[string]string{
	"windows": "windows",
	"linux":   "linux",
	"macos":   "mac",
}

var ARCH = map[string]string{
	"amd64": "x64",
	"386":   "x32",
	"arm64": "arm",
	"arm":   "arm",
}

func checkJRE() {
	//if checkJavaBin() {
	//	return
	//}
	if whether_download_java_or_not() {
		download_java()
	}
}

func checkJavaBin() bool {
	jpath, err := exec.LookPath(JAVA)
	if err != nil {
		return false
	}
	javaPath = jpath
	return true
}

func whether_download_java_or_not() bool {
	fmt.Print("未检测到Java环境，是否下载[y/n]：")
	return yes_or_not()
}

func whether_use_native_or_not() bool {
	fmt.Print("是否下载32位(Mirai Native)[y/n]：")
	return yes_or_not()
}

func yes_or_not() bool {
	var opt string
	for {
		fmt.Scanln(&opt)
		if opt == "y" || opt == "Y" {
			return true
		} else if opt == "n" || opt == "N" {
			return false
		}
		fmt.Println("输入格式不正确")
	}
}

func download_java() error {
	arch := runtime.GOARCH
	if arch == "windows" && whether_use_native_or_not() {
		arch = "386"
	}
	url := fmt.Sprintf(MIRROR, JAVA_VERSION, JDK_OR_JRE, ARCH[arch], OS[runtime.GOOS])
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return err
	}
	var links []string
	doc.Find("td.link>a").Each(func(i int, selection *goquery.Selection) {
		link, _ := selection.Attr("href")
		links = append(links, link)
	})

	return nil
}
