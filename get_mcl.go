package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"
)

const (
	GETVERSION = "https://api.github.com/repos/%s/releases/latest"
	REPOURL    = "iTXTech/mcl-installer"
)

var INDEXMAP = map[string]int{
	"linux":   0,
	"macos":   1,
	"windows": 2,
}

func get_url() (string, error) {
	resp, err := http.Get(fmt.Sprintf(GETVERSION, REPOURL))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("请求失败：%s", resp.Status)
	}
	data := make(map[string]interface{})
	json.NewDecoder(resp.Body).Decode(&data)
	log.Printf("target version: %s\t%s\n", data["tag_name"], data["published_at"])
	log.Println(data["body"])
	download_url := data["assets"].([]interface{})[INDEXMAP[runtime.GOOS]].(map[string]interface{})["browser_download_url"]
	return download_url, nil
}

func get_mcl() string{
	downUrl, err := get_url()
	if err != nil{
		log.Error(err)
	}
	return downUrl
}
