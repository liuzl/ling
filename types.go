package ling

import (
	"fmt"
)

type Token struct {
	Raw   string `json:"raw"`
	Start int    `json:"start"`
	End   int    `json:"end"`
}

func (t *Token) String() string {
	return fmt.Sprintf("%s[%d:%d]", t.Raw, t.Start, t.End)
}

type Tokenizer interface {
	Tokenize(text string) []*Token
}
