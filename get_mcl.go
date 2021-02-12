package main

import (
	//"log"
	"fmt"
)

const REPOURL = "iTXTech/mirai-console-loader"

func get_mcl() {
	downUrl, err := get_latest_version_url(REPOURL, 0)
	fmt.Println(downUrl)
	fmt.Println(err.Error())
}
