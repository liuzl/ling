package util

import (
	"fmt"
	"github.com/liuzl/da"
	"strings"
)

type is func(rune) bool

func StringIs(s string, f is) bool {
	for _, c := range s {
		if !f(c) {
			return false
		}
	}
	return true
}

func Convert(in string, dicts []*da.Dict) (string, error) {
	r := []rune(in)
	var tokens []string
	for i := 0; i < len(r); {
		s := r[i:]
		var token string
		max := 0
		for _, dict := range dicts {
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

func DictConvert(in []string, m map[string]string) ([]string, error) {
	if m == nil {
		return nil, fmt.Errorf("lemma dict is nil")
	}
	var ret []string
	for _, token := range in {
		if str, has := m[token]; has {
			ret = append(ret, str)
		} else {
			ret = append(ret, token)
		}
	}
	return ret, nil
}
