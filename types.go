package ling

import (
	"fmt"
)

type Document struct {
	Text   string   `json:"text"`
	Tokens []*Token `json:"tokens"`
	Spans  []*Span  `json:"spans"`
	Lang   string   `json:"lang"`
}

func (d *Document) String() string {
	return d.Text
}

type Token struct {
	Doc       *Document `json:"-"`
	Text      string    `json:"text"`
	StartByte int       `json:"start_byte"`
	EndByte   int       `json:"end_byte"`
	Lower     string    `json:"lower"`
	Norm      string    `json:"norm"`
	Stem      string    `json:"stem"`
}

func (t *Token) String() string {
	return fmt.Sprintf("(%s/%s/%s)[%d:%d]",
		t.Text, t.Norm, t.Stem, t.StartByte, t.EndByte)
}

type Span struct {
	Doc   *Document `json:"-"`
	Start int       `json:"start"`
	End   int       `json:"end"`
	Label string    `json:"label"`
	From  string    `json:"from"`
}
