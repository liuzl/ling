package ling

import (
	"fmt"
	"github.com/liuzl/segment"
	"strings"
	"unicode"
)

func Tokenize(text string) []string {
	var ret []string
	seg := segment.NewSegmenterDirect([]byte(text))
	for seg.Segment() {
		ret = append(ret, seg.Text())
	}
	return ret
}

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
	for i, word := range Tokenize(d.Text) {
		l := len([]byte(word))
		token := &Token{Doc: d, Text: word, Type: Type(word),
			I: i, StartByte: pos, EndByte: pos + l,
			Annotations: map[string]string{"lowerd": strings.ToLower(word)}}
		pos += l
		tokens = append(tokens, token)
	}
	d.Tokens = tokens
	return nil
}
