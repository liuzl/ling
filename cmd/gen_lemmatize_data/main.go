package main

import (
	"bytes"
	"fmt"
	"github.com/liuzl/ling/util"
	"log"
	"strconv"
	"strings"
)

func getInfo(lang string) (srcUrl, dstFile, varName string) {
	srcUrl = fmt.Sprintf(`https://raw.githubusercontent.com/unimorph/%s/master/%s`,
		lang, lang)
	dstFile = fmt.Sprintf(`src/github.com/liuzl/ling/lemmatize/%s.go`, lang)
	varName = strings.ToUpper(lang)
	return
}

func gen(lang string) {
	srcUrl, dstFile, varName := getInfo(lang)
	log.Println("Fetching " + srcUrl)
	c := util.UrlToZipContent(srcUrl)
	output := bytes.Buffer{}
	output.WriteString("package lemmatize\n\n")
	output.WriteString(fmt.Sprintf("var %s = %s\n", varName, strconv.Quote(c)))
	output.WriteString(fmt.Sprintf("\nfunc init() {\n\tgenDict(%s, %s)\n}\n",
		strconv.Quote(lang), varName))
	log.Println("Writing new " + dstFile)
	util.WriteFile(dstFile, output.Bytes())
}

func main() {
	for _, lang := range langs {
		gen(lang)
	}
}
