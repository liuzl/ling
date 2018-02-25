package lemma

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/golang/glog"
	"github.com/liuzl/da"
	"github.com/liuzl/tokenizer"
	"strings"
)

var cmnDicts []*da.Dict

func initDict(contents ...*string) error {
	for _, c := range contents {
		data, err := base64.StdEncoding.DecodeString(*c)
		if err != nil {
			glog.Error(err)
			return err
		}
		reader, err := gzip.NewReader(bytes.NewReader(data))
		if err != nil {
			glog.Error(err)
			return err
		}
		dict, err := da.Build(reader)
		if err != nil {
			glog.Error(err)
			return err
		}
		cmnDicts = append(cmnDicts, dict)
	}
	Funcs["cmn"] = func(in []string) ([]string, error) {
		inStr := strings.Join(in, "")
		out, err := Convert(inStr)
		if err != nil {
			return nil, err
		}
		ret := tokenizer.Tokenize(out)
		if len(ret) == len(in) {
			return ret, nil
		} else {
			return nil, fmt.Errorf("len(ret)!=len(in) of cmn lemma function")
		}
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
