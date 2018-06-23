package ling

import (
	"fmt"

	"github.com/liuzl/unidecode"
)

const Unidecode = "unidecode"

func init() {
	Processors[Unidecode] = &Unidecoder{}
}

type Unidecoder struct {
}

func (u *Unidecoder) Process(d *Document) error {
	if d == nil || len(d.Text) == 0 {
		return nil
	}
	if len(d.Tokens) == 0 {
		return fmt.Errorf("tokenization is required for unidecode")
	}
	for _, token := range d.Tokens {
		var s string
		if s = token.Annotations[Norm]; s == "" {
			s = token.Text
		}
		token.Annotations[Unidecode] = unidecode.Unidecode(s)
	}
	return nil
}
