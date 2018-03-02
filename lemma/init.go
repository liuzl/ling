package lemma

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"github.com/golang/glog"
	"github.com/liuzl/ling/util"
	"io/ioutil"
	"strings"
)

var Funcs = make(map[string]util.ConvertFunc)

func genDict(lang, body string) error {
	data, err := base64.StdEncoding.DecodeString(body)
	if err != nil {
		glog.Error(err)
		return err
	}
	reader, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		glog.Error(err)
		return err
	}
	rawBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		glog.Error(err)
		return err
	}
	m := make(map[string]string)
	for _, line := range strings.Split(string(rawBytes), "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}
		items := strings.Fields(line)
		if len(items) != 3 || items[0] == items[1] || len(items[1]) < 3 {
			// some short words have problems
			continue
		}
		m[items[1]] = items[0]
	}
	Funcs[lang] = func(in []string) ([]string, error) {
		return util.DictConvert(in, m)
	}
	return nil
}
