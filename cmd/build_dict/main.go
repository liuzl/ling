package main

import (
	"flag"
	"log"

	"github.com/liuzl/da"
)

var (
	i = flag.String("i", "input.txt", "input dict file")
	o = flag.String("o", "dict", "output dict")
)

func main() {
	flag.Parse()
	d, err := da.BuildFromFile(*i)
	if err != nil {
		log.Fatal(err)
	}
	if err = d.Save(*o); err != nil {
		log.Fatal(err)
	}
	log.Printf("dict generated to %s\n", *o)
}
