package ling

import (
	"flag"
	"fmt"

	"github.com/liuzl/d"
)

var (
	dictDir = flag.String("dict_dir", "dict", "dictionary dir")
	dictWeb = flag.Bool("dict_web", false, "dictionary web api flag")
)

type DictTagger struct {
	*d.Dictionary
}

func NewDictTagger() (*DictTagger, error) {
	dict, err := d.Load(*dictDir)
	if err != nil {
		return nil, err
	}
	if *dictWeb {
		dict.RegisterWeb()
	}
	return &DictTagger{dict}, nil
}

func (t *DictTagger) Process(d *Document) error {
	if d == nil || len(d.Text) == 0 {
		return nil
	}
	if len(d.Tokens) == 0 {
		return fmt.Errorf("tokenization required")
	}
	r := []rune(d.Text)
	for i := 0; i < len(r); i++ {
		startByte := len(string(r[:i]))
		ret, err := t.PrefixMatch(string(r[i:]))
		if err != nil {
			return err
		}
		for k, v := range ret {
			if len(v) < 1 {
				continue
			}
			start := -1
			end := -1
			for _, token := range d.Tokens {
				if token.StartByte == startByte {
					start = token.I
				}
				if token.EndByte == startByte+len(k) {
					end = token.I + 1
				}
			}
			if start == -1 || end == -1 {
				continue
			}
			span := &Span{Doc: d, Start: start, End: end,
				Annotations: map[string]interface{}{"from": "dict", "value": v}}
			d.Spans = append(d.Spans, span)
		}
	}
	return nil
}
