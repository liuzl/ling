package ling

import (
	"fmt"

	"github.com/liuzl/ling/lemmatize"
)

// Lemma processor name
const Lemma = "lemma"

func init() {
	Processors[Lemma] = &Lemmatizer{}
}

// Lemmatizer is the processor for lemmatization
type Lemmatizer struct {
}

// Process is the function to annotate documents
func (l *Lemmatizer) Process(d *Document) error {
	if d == nil || len(d.Text) == 0 {
		return nil
	}
	if len(d.Tokens) == 0 {
		return fmt.Errorf("tokenization required")
	}

	for _, lang := range d.Langs {
		if f, has := lemmatize.Funcs[lang]; has {
			ret, err := f(d.XTokens(Norm))
			if err != nil {
				return err
			}
			if len(ret) != len(d.Tokens) {
				continue
			}
			for i, str := range ret {
				d.Tokens[i].Annotations[Lemma] = str
			}
		}
	}
	return nil
}
