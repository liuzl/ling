package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func getInfo(lang string) (srcUrl, dstFile, varName string) {
	srcUrl = fmt.Sprintf(`https://raw.githubusercontent.com/unimorph/%s/master/%s`,
		lang, lang)
	dstFile = fmt.Sprintf(`src/github.com/liuzl/ling/resources/lemma/%s.go`, lang)
	varName = strings.ToUpper(lang)
	return
}

func fetchURL(url string) []byte {
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

func writeFile(filePath string, data []byte) {
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

func gen(lang string) {
	srcUrl, dstFile, varName := getInfo(lang)
	log.Println("Fetching " + srcUrl)
	body := fetchURL(srcUrl)
	var compressed bytes.Buffer
	w := gzip.NewWriter(&compressed)
	w.Write(body)
	w.Close()
	c := base64.StdEncoding.EncodeToString(compressed.Bytes())
	output := bytes.Buffer{}
	output.WriteString("package lemma\n\n")
	output.WriteString(fmt.Sprintf("var %s string = %s\n", varName, strconv.Quote(c)))
	output.WriteString(fmt.Sprintf("func init() {\n\tgenDict(%s, %s)\n}",
		strconv.Quote(lang), varName))
	log.Println("Writing new " + dstFile)
	writeFile(dstFile, output.Bytes())
	//writeFile("src/github.com/liuzl/ling/resources/lemma/"+lang+".txt", body)
}

func main() {
	for _, lang := range langs {
		gen(lang)
	}
}
