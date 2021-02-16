package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/mholt/archiver"
)

type WriteCounter struct {
	Total  uint64
	Length uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.printProgress()
	return n, nil
}

func (wc WriteCounter) printProgress() {
	fmt.Printf("\r%s", strings.Repeat(" ", 35))
	fmt.Printf("\r下载中...\t[%s/%s]", humanize.Bytes(wc.Total), humanize.Bytes(wc.Length))
}

func downloadFile(filepath string, url string) error {
	out, err := os.Create(filepath + ".tmp")
	if err != nil {
		return err
	}
	resp, err := http.Get(url)
	if err != nil {
		out.Close()
		return err
	}
	defer resp.Body.Close()
	counter := &WriteCounter{0, uint64(resp.ContentLength)}
	if _, err = io.Copy(out, io.TeeReader(resp.Body, counter)); err != nil {
		out.Close()
		return err
	}
	fmt.Print("\n")
	out.Close()
	if err = os.Rename(filepath+".tmp", filepath); err != nil {
		return err
	}
	return nil
}

func unpack(origin string, target string) error {
	err := archiver.Unarchive(origin, target)
	if err != nil {
		return err
	}
	os.Remove(origin)
	return nil
}
