package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
)

func save(br *bufio.Reader, fname string) bool {
	if br == nil {
		return false
	}
	f, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}
	_, err = br.WriteTo(f)
	if err != nil {
		panic(err)
	}
	_ = f.Close()
	log.Println("下载成功")
	return true
}

func downFile(url string) *bufio.Reader {
	req, err := http.NewRequest("GET", url, nil)
	log.Println("开始下载...")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalln()
	}
	return bufio.NewReaderSize(resp.Body, 1<<20)
}
