package environment

import (
	"fmt"
	"path/filepath"
	"regexp"
	"runtime"
	"sync"

	down "github.com/EricTianC/GraiaOK/download"
	"github.com/PuerkitoBio/goquery"
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

func (es *EnvSpace) CheckJava(js chan<- bool) {
	if jenv.IsInstalled() { //之所以把IsInstalled单独列成函数，是因为有的时候需要优先本地已装好的环境，有的时候又要优先自动装的环境
		js <- true
		return
	}

	//检查EnvSpace中是否有Java可执行文件
	if jenv.LookForExecFileInSpace(es) {
		js <- true
		return
	}
	js <- false
}

func (es *EnvSpace) DownloadJava(gwg *sync.WaitGroup, complete chan<- struct{}) error {
	defer gwg.Done()
	//下载Java

	//确认是否使用32位
	arch := runtime.GOARCH
	if runtime.GOOS == "windows" && yesorNot("是否下载三十二位版本", false) {
		arch = "386"
	}
	//解析镜像列表Html
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
	//正则筛选符合的文件
	rt := fmt.Sprintf("^OpenJDK%sU-%s_%s_%s_hotspot_[0-9]{1,2}\\.[0-9]{1,2}\\.[0-9]{1,2}_[0-9]\\.(zip|tar\\.gz)", JAVA_VERSION, JDK_OR_JRE, ARCH[arch], OS[runtime.GOOS])
	r, _ := regexp.Compile(rt)
	var arch_url, name string
	for _, link := range links {
		if r.MatchString(link) {
			name = link
			arch_url = url + link
			break
		}
	}
	down.DownloadFile(name, arch_url, "下载Java")
	down.Unpack(name, filepath.Join(es.BasePath, jenv.BasePath))
	go func() { complete <- struct{}{} }()
	if jenv.LookForExecFileInSpace(es) {
		return nil
	}
	return fmt.Errorf("无法配置Java环境")
}
