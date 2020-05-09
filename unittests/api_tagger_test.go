package unittests

import (
	"testing"

	"github.com/juju/errors"
	"github.com/liuzl/ling"
)

func TestAPITagger(t *testing.T) {
	cases := []string{
		`天津大学离北京也不太远，电话是02227890949，主页是https://www.tju.edu.cn，计算机`,
		`https://crawler.club是爬虫主页`,
		`刘晓明是个好同志，他和周晓华是一对儿好朋友`,
	}

	nlp, err := ling.DefaultNLP()
	if err != nil {
		t.Error(err)
	}
	tagger, err := ling.NewAPITagger("http://127.0.0.1:5002/api")
	if err != nil {
		t.Error(err)
	}
	if err = nlp.AddTagger(tagger); err != nil {
		t.Error(err)
	}
	for _, c := range cases {
		d := ling.NewDocument(c)
		if err = nlp.Annotate(d); err != nil {
			t.Error(errors.ErrorStack(err))
		}
		t.Logf("spans:%s", d.Spans)
		for i, s := range d.Spans {
			t.Logf("span %d %s: %+v", i, s, s.Annotations)
		}
	}
}
