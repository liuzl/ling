package tagger

import (
	"fmt"
	"github.com/liuzl/ling"
)

type RegexTagger struct {
}

func (t *RegexTagger) Process(d *ling.Document) error {
	if d == nil || len(d.Text) == 0 {
		return fmt.Errorf("document is empty")
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
			span := &ling.Span{Doc: d, Start: start, End: end,
				Annotations: map[string]string{"type": typ}}
			d.Spans = append(d.Spans, span)
		}
	}
	return nil
}
