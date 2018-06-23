package ling

import (
	"strings"
	"unicode"

	"github.com/liuzl/ling/util"
	"github.com/liuzl/tokenizer"
)

const Lower = "lower"

func init() {
	Processors["_"] = &Tokenizer{}
}

func Type(text string) TokenType {
	switch {
	case util.StringIs(text, unicode.IsSpace):
		return Space
	case util.StringIs(text, unicode.IsSymbol):
		return Symbol
	case util.StringIs(text, unicode.IsNumber):
		return Number
	case util.StringIs(text, unicode.IsPunct):
		return Punct
	case util.StringIs(text, func(r rune) bool {
		return unicode.Is(unicode.Scripts["Han"], r)
	}):
		return Han
	}
	return Word
}

type Tokenizer struct {
}

func (t *Tokenizer) Process(d *Document) error {
	if d == nil || len(d.Text) == 0 {
		return nil
	}
	var tokens []*Token
	var pos int = 0
	for i, item := range tokenizer.TokenizePro(d.Text) {
		word := item.Text
		l := len([]byte(word))
		token := &Token{Doc: d, Text: word, Type: Type(word),
			I: i, StartByte: pos, EndByte: pos + l,
			Annotations: map[string]string{Lower: strings.ToLower(word)}}
		if item.Norm != "" {
			token.Annotations[Norm] = item.Norm
		}
		pos += l
		tokens = append(tokens, token)
	}
	d.Tokens = tokens
	return nil
}
