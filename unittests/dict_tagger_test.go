package unittests

import (
	"testing"

	"github.com/liuzl/ling"
)

func TestDictTagger(t *testing.T) {
	cases := []string{
		`天津大学离北京也不太远`,
		`https://crawler.club是爬虫主页`,
	}

	nlp, err := ling.DefaultNLP()
	if err != nil {
		t.Error(err)
	}
	tagger, err := ling.NewDictTagger("dict")
	if err != nil {
		t.Error(err)
	}
	if err = nlp.AddTagger(tagger); err != nil {
		t.Error(err)
	}
	for _, c := range cases {
		d := ling.NewDocument(c)
		if err = nlp.Annotate(d); err != nil {
			t.Error(err)
		}
		t.Logf("spans:%s", d.Spans)
		for i, s := range d.Spans {
			t.Logf("span %d %s: %+v", i, s, s.Annotations)
		}
	}
}
