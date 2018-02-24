package ling

import (
	"bytes"
	"github.com/liuzl/franco"
)

func NewDocument(text string) *Document {
	d := &Document{Text: text}
	r := franco.Detect(text)
	if len(r) > 0 {
		if len(r) <= 10 || r[0].Code == "eng" || r[0].Code == "rus" {
			d.Lang = r[0].Code
		}
		for _, lang := range r {
			d.Langs = append(d.Langs, lang.Code)
		}
	}
	return d
}

func (d *Document) String() string {
	return d.Text
}

func (d *Document) Norm() string {
	var buffer bytes.Buffer
	for _, token := range d.Tokens {
		buffer.WriteString(token.Annotations["norm"])
	}
	return buffer.String()
}
