package main

import (
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	fetch("http://cbt.bg")
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)

	defer f.Close()

	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	return local, n, err
}
