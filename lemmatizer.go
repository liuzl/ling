package ling

import (
	"fmt"
)

type Lemmatizer struct {
}

func (self *Lemmatizer) Process(d *Document) error {
	if d == nil || len(d.Text) == 0 {
		return fmt.Errorf("ducument is empty")
	}
	if len(d.Tokens) == 0 {
		return fmt.Errorf("tokenization required")
	}
	// TODO
	return nil
}
