package util

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func FetchURL(url string) []byte {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		log.Fatalf("Error fetching URL '%s': %s", url, err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading body: %s", err)
	}
	return body
}

func UrlToZipContent(srcUrl string) string {
	body := FetchURL(srcUrl)
	var compressed bytes.Buffer
	w := gzip.NewWriter(&compressed)
	w.Write(body)
	w.Close()
	return base64.StdEncoding.EncodeToString(compressed.Bytes())
}

func WriteFile(filePath string, data []byte) {
	gopath, found := os.LookupEnv("GOPATH")
	if !found {
		log.Fatal("Missing $GOPATH environment variable")
	}
	path := filepath.Join(gopath, filePath)
	err := ioutil.WriteFile(path, data, os.FileMode(0664))
	if err != nil {
		log.Fatalf("Error writing '%s': %s", path, err)
	}
}
