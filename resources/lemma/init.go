package lemma

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/golang/glog"
	"io/ioutil"
	"strings"
)

type lemmaFunc func([]string) ([]string, error)

var Funcs map[string]lemmaFunc = make(map[string]lemmaFunc)
var Dict map[string]map[string]string = make(map[string]map[string]string)

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
	Dict[lang] = m
	Funcs[lang] = func(in []string) ([]string, error) {
		return dictMatch(in, m)
	}
	return nil
}

func dictMatch(in []string, m map[string]string) ([]string, error) {
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
