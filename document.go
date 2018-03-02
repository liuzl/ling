package ling

import (
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

func (d *Document) NewSpan(start, end int) *Span {
	tokenCnt := len(d.Tokens)
	if tokenCnt == 0 {
		return nil
	}
	if start < 0 {
		start = 0
	}
	if end > tokenCnt {
		end = tokenCnt
	}
	return &Span{Doc: d, Start: start, End: end}
}

func (d *Document) XTokens(anno string) []string {
	var ret []string
	for _, token := range d.Tokens {
		t, has := token.Annotations[anno]
		if !has {
			t = token.Annotations["norm"]
		}
		ret = append(ret, t)
	}
	return ret
}
