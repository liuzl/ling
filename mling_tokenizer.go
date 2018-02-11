package ling

import (
	"github.com/liuzl/go-porterstemmer"
	"github.com/liuzl/segment"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"strings"
	"unicode"
)

type MlingTokenizer struct {
	trans transform.Transformer
	err   error
}

func NewMlingTokenizer() *MlingTokenizer {
	tok := &MlingTokenizer{}
	isMn := func(r rune) bool {
		return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
	}
	tok.trans = transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	return tok
}

func (self *MlingTokenizer) Tokenize(text string) []*Token {
	var ret []*Token
	seg := segment.NewSegmenterDirect([]byte(text))
	var pos int = 0
	for seg.Segment() {
		l := len(seg.Bytes())
		token := &Token{Text: seg.Text(), StartByte: pos, EndByte: pos + l}
		token.Lower = strings.ToLower(token.Text)
		res, _, err := transform.String(self.trans, token.Lower)
		if err != nil {
			self.err = err
		} else {
			token.Norm = res
		}
		token.Stem = porterstemmer.StemString(token.Norm)
		pos += l
		ret = append(ret, token)
	}
	return ret
}
