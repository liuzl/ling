package ling

import (
	"fmt"
)

func init() {
	Processors["regex"] = &RegexTagger{}
}

type RegexTagger struct {
}

func (t *RegexTagger) Process(d *Document) error {
	if d == nil || len(d.Text) == 0 {
		return nil
	}
	if len(d.Tokens) == 0 {
		return fmt.Errorf("tokenization required")
	}
	for typ, re := range Regexes {
		matches := re.FindAllStringIndex(d.Text, -1)
		for _, match := range matches {
			start := -1
			end := -1
			for _, token := range d.Tokens {
				if token.StartByte == match[0] {
					start = token.I
				}
				if token.EndByte == match[1] {
					end = token.I + 1
				}
			}
			if start == -1 || end == -1 {
				continue
			}
			span := &Span{Doc: d, Start: start, End: end,
				Annotations: map[string]interface{}{
					"from": "regex", "type": typ}}
			d.Spans = append(d.Spans, span)
		}
	}
	return nil
}
