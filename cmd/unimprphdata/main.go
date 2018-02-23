package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	t = flag.String("t", "all", "unimorph, chinese, all")
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

func urlToContent(srcUrl string) string {
	body := fetchURL(srcUrl)
	var compressed bytes.Buffer
	w := gzip.NewWriter(&compressed)
	w.Write(body)
	w.Close()
	return base64.StdEncoding.EncodeToString(compressed.Bytes())
}

func gen(lang string) {
	srcUrl, dstFile, varName := getInfo(lang)
	log.Println("Fetching " + srcUrl)
	c := urlToContent(srcUrl)
	output := bytes.Buffer{}
	output.WriteString("package lemma\n\n")
	output.WriteString(fmt.Sprintf("var %s string = %s\n", varName, strconv.Quote(c)))
	output.WriteString(fmt.Sprintf("func init() {\n\tgenDict(%s, %s)\n}",
		strconv.Quote(lang), varName))
	log.Println("Writing new " + dstFile)
	writeFile(dstFile, output.Bytes())
	//writeFile("src/github.com/liuzl/ling/resources/lemma/"+lang+".txt", body)
}

func chinese() {
	srcUrl1 := `https://raw.githubusercontent.com/liuzl/gocc/master/dictionary/TSPhrases.txt`
	srcUrl2 := `https://raw.githubusercontent.com/liuzl/gocc/master/dictionary/TSCharacters.txt`
	dstFile := `src/github.com/liuzl/ling/resources/lemma/chinese.go`
	log.Println("Fetching " + srcUrl1)
	src1 := urlToContent(srcUrl1)
	log.Println("Fetching " + srcUrl2)
	src2 := urlToContent(srcUrl2)

	output := bytes.Buffer{}
	output.WriteString("package lemma\n\n")
	output.WriteString(fmt.Sprintf("var cmnPhrases string = %s\n", strconv.Quote(src1)))
	output.WriteString(fmt.Sprintf("var cmnCharacters string = %s\n", strconv.Quote(src2)))
	output.WriteString(fmt.Sprintf("func init() {\n\tinitDict(&cmnPhrases, &cmnCharacters)\n}"))
	log.Println("Writing new " + dstFile)
	writeFile(dstFile, output.Bytes())
}

func main() {
	flag.Parse()
	switch *t {
	case "unimorph":
		for _, lang := range langs {
			gen(lang)
		}
	case "chinese":
		chinese()
	case "all":
		for _, lang := range langs {
			gen(lang)
		}
		chinese()
	default:
		flag.PrintDefaults()
	}
}
