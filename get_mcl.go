package main

import (
	"log"
	"fmt"
)

const REPOURL = "iTXTech/mirai-console-loader"

func get_mcl() {
	downUrl, err := get_latest_version_url(REPOURL, 0)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(downUrl)
}
