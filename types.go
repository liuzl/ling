package ling

import (
	"bytes"
)

type Document struct {
	Text   string   `json:"text"`
	Tokens []*Token `json:"tokens"`
	Spans  []*Span  `json:"spans"`
	Lang   string   `json:"lang"`
	Langs  []string `json:"langs"`
}

type TokenType byte

//go:generate jsonenums -type=TokenType
//go:generate stringer -type=TokenType
const (
	EOF TokenType = iota
	Space
	Symbol
	Number
	Letters
	Punct
	Word
)

type Token struct {
	Doc         *Document         `json:"-"`
	Text        string            `json:"text"`
	Type        TokenType         `json:"type"`
	Script      string            `json:"script"`
	I           int               `json:"i"`
	StartByte   int               `json:"start_byte"`
	EndByte     int               `json:"end_byte"`
	Annotations map[string]string `json:"annotations"`
}

func (t *Token) String() string {
	return t.Text
}

type Span struct {
	Doc         *Document              `json:"-"`
	Start       int                    `json:"start"`
	End         int                    `json:"end"`
	Annotations map[string]interface{} `json:"annotations"`
}

func (s *Span) String() string {
	output := bytes.Buffer{}
	for i := s.Start; i < s.End; i++ {
		output.WriteString(s.Doc.Tokens[i].String())
	}
	return output.String()
}

type Processor interface {
	Process(d *Document) error
}

var Processors = make(map[string]Processor)
