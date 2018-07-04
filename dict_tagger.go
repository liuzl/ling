package ling

import (
	"fmt"

	"github.com/liuzl/da"
)

type DictTagger struct {
	dict *da.Dict
}

func NewDictTagger(dir string) (*DictTagger, error) {
	dict, err := da.Load(dir)
	if err != nil {
		return nil, err
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
		ret, err := t.dict.PrefixMatch(string(r[i:]))
		if err != nil {
			return err
		}
		for k, v := range ret {
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
				Annotations: map[string]interface{}{
					"from": "dict", "value": v}}
			d.Spans = append(d.Spans, span)
		}
	}
	return nil
}
