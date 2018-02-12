package ling

import (
	"fmt"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"golang.org/x/text/width"
	"strings"
	"unicode"
)

var trans transform.Transformer = transform.Chain(
	norm.NFD,
	transform.RemoveFunc(func(r rune) bool {
		return unicode.Is(unicode.Mn, r)
	}),
	norm.NFC)

var replacer *strings.Replacer = strings.NewReplacer(
	`｡`, `.`, // half period in Chinese
	`。`, `.`, // full period in Chinese
	`【`, `[`,
	`】`, `]`,
	`“`, `"`,
	`”`, `"`,
	`‘`, `'`,
	`’`, `'`,
)

type Normalizer struct {
}

func (self *Normalizer) Process(d *Document) error {
	if d == nil || len(d.Text) == 0 {
		return fmt.Errorf("document is empty")
	}
	if len(d.Tokens) == 0 {
		return fmt.Errorf("tokenization required")
	}
	for _, token := range d.Tokens {
		res, _, err := transform.String(trans, token.Annotations["lowerd"])
		if err != nil {
			return err
		}
		// full to half
		token.Annotations["norm"] = replacer.Replace(width.Narrow.String(res))
	}
	return nil
}
