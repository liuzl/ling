package normalize

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/golang/glog"
	"github.com/liuzl/da"
	"github.com/liuzl/ling/util"
	"github.com/liuzl/tokenizer"
	"strings"
)

var Funcs = make(map[string]util.ConvertFunc)

func genFuncs(lang string, contents ...*string) error {
	var dicts []*da.Dict
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
		dicts = append(dicts, dict)
	}
	Funcs[lang] = func(in []string) ([]string, error) {
		out, err := util.Convert(strings.Join(in, ""), dicts)
		if err != nil {
			return nil, err
		}
		ret := tokenizer.Tokenize(out)
		if len(ret) == len(in) {
			return ret, nil
		}
		return nil, fmt.Errorf("len(ret):%d!=len(in):%d of %s norm function",
			len(ret), len(in), lang)
	}
	return nil
}
