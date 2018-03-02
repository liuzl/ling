package main

import (
	"bytes"
	"fmt"
	"github.com/liuzl/ling/util"
	"log"
	"strconv"
)

func chinese() {
	lang := `cmn`
	srcUrl1 := `https://raw.githubusercontent.com/liuzl/gocc/master/dictionary/TSPhrases.txt`
	srcUrl2 := `https://raw.githubusercontent.com/liuzl/gocc/master/dictionary/TSCharacters.txt`
	dstFile := fmt.Sprintf(`src/github.com/liuzl/ling/normalize/%s.go`, lang)
	log.Println("Fetching " + srcUrl1)
	src1 := util.UrlToZipContent(srcUrl1)
	log.Println("Fetching " + srcUrl2)
	src2 := util.UrlToZipContent(srcUrl2)

	output := bytes.Buffer{}
	output.WriteString("package normalize\n\n")
	output.WriteString(fmt.Sprintf("var cmnPhrases = %s\n", strconv.Quote(src1)))
	output.WriteString(fmt.Sprintf("var cmnCharacters = %s\n", strconv.Quote(src2)))
	output.WriteString(fmt.Sprintf("\nfunc init() {\n\tgenFuncs(%s, &cmnPhrases, &cmnCharacters)\n}\n", strconv.Quote(lang)))
	log.Println("Writing new " + dstFile)
	util.WriteFile(dstFile, output.Bytes())
}

func main() {
	chinese()
}
