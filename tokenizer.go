package ling

import (
	"fmt"
	"github.com/liuzl/tokenizer"
	"strings"
	"unicode"
)

func Type(text string) TokenType {
	switch {
	case StringIs(text, unicode.IsSpace):
		return Space
	case StringIs(text, unicode.IsSymbol):
		return Symbol
	case StringIs(text, unicode.IsNumber):
		return Number
	case StringIs(text, unicode.IsPunct):
		return Punct
	case StringIs(text, func(r rune) bool {
		return unicode.Is(unicode.Scripts["Han"], r)
	}):
		return Han
	}
	return Word
}

type Tokenizer struct {
}

func (self *Tokenizer) Process(d *Document) error {
	if d == nil || len(d.Text) == 0 {
		return fmt.Errorf("document is empty")
	}
	var tokens []*Token
	var pos int = 0
	for i, item := range tokenizer.TokenizePro(d.Text) {
		word := item.Text
		l := len([]byte(word))
		token := &Token{Doc: d, Text: word, Type: Type(word),
			I: i, StartByte: pos, EndByte: pos + l,
			Annotations: map[string]string{"lowerd": strings.ToLower(word)}}
		if item.Norm != "" {
			token.Annotations["norm"] = item.Norm
		}
		pos += l
		tokens = append(tokens, token)
	}
	d.Tokens = tokens
	return nil
}
