package ling

import (
	"fmt"

	"github.com/liuzl/ling/normalize"
)

const Norm = "norm"

func init() {
	Processors[Norm] = &Normalizer{}
}

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
