package lemma

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"github.com/liuzl/da"
	"strings"
)

var cmnDicts []*da.Dict

func initDict(contents ...*string) error {
	for _, c := range contents {
		data, err := base64.StdEncoding.DecodeString(*c)
		if err != nil {
			return err
		}
		reader, err := gzip.NewReader(bytes.NewReader(data))
		if err != nil {
			return err
		}
		dict, err := da.Build(reader)
		if err != nil {
			return err
		}
		cmnDicts = append(cmnDicts, dict)
	}
	return nil
}

func Convert(in string) (string, error) {
	r := []rune(in)
	var tokens []string
	for i := 0; i < len(r); {
		s := r[i:]
		var token string
		max := 0
		for _, dict := range cmnDicts {
			ret, err := dict.PrefixMatch(string(s))
			if err != nil {
				return "", err
			}
			if len(ret) > 0 {
				o := ""
				for k, v := range ret {
					if len(k) > max {
						max = len(k)
						token = v[0]
						o = k
					}
				}
				i += len([]rune(o))
				break
			}
		}
		if max == 0 { //no match
			token = string(r[i])
			i++
		}
		tokens = append(tokens, token)
	}
	return strings.Join(tokens, ""), nil
}
