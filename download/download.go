package download

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/mholt/archiver"
	"github.com/qianlnk/pgbar"
)

type WriteCounter struct {
	Bar *pgbar.Bar
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Bar.Add(int(n))
	return n, nil
}

func DownloadFile(filepath string, url string) error {
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
	bar := pgbar.NewBar(0, "下载中", int(resp.ContentLength))
	counter := &WriteCounter{bar}
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

func Unpack(origin string, target string) error {
	err := archiver.Unarchive(origin, target)
	if err != nil {
		return err
	}
	os.Remove(origin)
	return nil
}
