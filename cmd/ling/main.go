package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/liuzl/ling"
)

var (
	input  = flag.String("input", "", "file of original text to read")
	output = flag.String("output", "", "file of tokenized text to write")
	typ    = flag.String("type", "token", "type: token, span")
)

func main() {
	flag.Parse()
	var err error
	var in, out *os.File //io.Reader
	var nlp *ling.Pipeline
	if *input == "" {
		in = os.Stdin
	} else {
		in, err = os.Open(*input)
		if err != nil {
			log.Fatal(err)
		}
		defer in.Close()
	}
	br := bufio.NewReader(in)

	if *output == "" {
		out = os.Stdout
	} else {
		out, err = os.OpenFile(*output, os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
	}
	if nlp, err = ling.DefaultNLP(); err != nil {
		log.Fatal(err)
	}
	for {
		line, c := br.ReadString('\n')
		if c == io.EOF {
			break
		}
		if c != nil {
			log.Fatal(c)
		}
		d := ling.NewDocument(strings.TrimSpace(line))
		if err = nlp.Annotate(d); err != nil {
			log.Fatal(err)
		}
		var ret []string
		if *typ == "span" {
			for _, s := range d.Spans {
				ret = append(ret, s.String())
			}
		} else {
			for _, t := range d.Tokens {
				if t.Type == ling.Space {
					continue
				}
				ret = append(ret, t.String())
			}
		}
		fmt.Fprintf(out, "%s\n", strings.Join(ret, " "))
	}
}
