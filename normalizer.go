package ling

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/liuzl/ling/normalize"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"golang.org/x/text/width"
)

const Norm = "norm"

func init() {
	Processors[Norm] = &Normalizer{}
}

var trans = transform.Chain(
	norm.NFD,
	transform.RemoveFunc(func(r rune) bool {
		return unicode.Is(unicode.Mn, r)
	}),
	norm.NFC)

var replacer = strings.NewReplacer(
	`｡`, `.`, // half period in Chinese
	`。`, `.`, // full period in Chinese
	`【`, `[`,
	`】`, `]`,
	`“`, `"`,
	`”`, `"`,
	`‘`, `'`,
	`’`, `'`,
	`—`, `-`,
	`〔`, `{`,
	`〕`, `}`,
	`《`, `<`,
	`》`, `>`,
)

// Normalizer is the processor for token normalization
type Normalizer struct {
}

// Process normalizes the tokens of Document d
func (n *Normalizer) Process(d *Document) error {
	if d == nil || len(d.Text) == 0 {
		return nil
	}
	if len(d.Tokens) == 0 {
		return fmt.Errorf("tokenization required")
	}
	for _, token := range d.Tokens {
		if _, has := token.Annotations[Norm]; has {
			continue
		}
		res, _, err := transform.String(trans, token.Annotations[Lower])
		if err != nil {
			return err
		}
		// full to half
		token.Annotations[Norm] = replacer.Replace(width.Narrow.String(res))
	}
	if f, has := normalize.Funcs[d.Lang]; has {
		ret, err := f(d.XTokens(Norm))
		if err != nil {
			return err
		}
		if len(ret) == len(d.Tokens) {
			for i, str := range ret {
				d.Tokens[i].Annotations[Norm] = str
			}
		}
	}
	return nil
}
