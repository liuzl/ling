package ling

import (
	"fmt"
)

type Document struct {
	Text   string   `json:"text"`
	Tokens []*Token `json:"tokens"`
	Spans  []*Span  `json:"spans"`
	Lang   string   `json:"lang"`
	Langs  []string `json:"langs"`
}

func (self *Document) NormTokens() []string {
	var ret []string
	for _, token := range self.Tokens {
		t, has := token.Annotations["norm"]
		if !has {
			t = token.Text
		}
		ret = append(ret, t)
	}
	return ret
}

type TokenType byte

const (
	EOF TokenType = iota
	Space
	Symbol
	Number
	Punct
	Han
	Word
)

type Token struct {
	Doc         *Document         `json:"-"`
	Text        string            `json:"text"`
	Type        TokenType         `json:"type"`
	I           int               `json:"i"`
	StartByte   int               `json:"start_byte"`
	EndByte     int               `json:"end_byte"`
	Annotations map[string]string `json:"annotations"`
}

func (t *Token) String() string {
	return fmt.Sprintf("(%q/%v){%+v}[%d:%d]",
		t.Text, t.Type, t.Annotations, t.StartByte, t.EndByte)
}

type Span struct {
	Doc         *Document         `json:"-"`
	Start       int               `json:"start"`
	End         int               `json:"end"`
	Annotations map[string]string `json:"annotations"`
}
