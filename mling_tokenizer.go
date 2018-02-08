package ling

import (
	"github.com/liuzl/segment"
)

type Option struct {
}

type MlingTokenizer struct {
	option Option
}

func NewMlingTokenizer() *MlingTokenizer {
	return &MlingTokenizer{}
}

func (self *MlingTokenizer) Tokenize(text string) []*Token {
	var ret []*Token
	seg := segment.NewSegmenterDirect([]byte(text))
	var pos int = 0
	for seg.Segment() {
		l := len(seg.Bytes())
		token := &Token{Raw: seg.Text(), Start: pos, End: pos + l}
		pos += l
		ret = append(ret, token)
	}
	return ret
}
