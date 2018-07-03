package unittests

import (
	"testing"

	"github.com/liuzl/ling"
)

func TestDictTagger(t *testing.T) {
	cases := []string{
		`一丝不挂的一分钟，是一前一后，有点意思吧`,
	}

	nlp, err := ling.DefaultNLP()
	if err != nil {
		t.Error(err)
	}
	tagger, err := ling.NewDictTagger("cedar")
	if err != nil {
		t.Error(err)
	}
	for _, c := range cases {
		d := ling.NewDocument(c)
		if err = nlp.Annotate(d); err != nil {
			t.Error(err)
		}
		if err = tagger.Process(d); err != nil {
			t.Error(err)
		}
		t.Logf("lang  :%s", d.Lang)
		t.Logf("langs :%s", d.Langs)
		t.Logf("tokens:%s", d.Tokens)
		t.Logf("spans:%s", d.Spans)
		t.Logf("lower :%s", d.XTokens(ling.Lower))
		t.Logf("norm  :%s", d.XTokens(ling.Norm))
		t.Logf("lemma :%s", d.XTokens(ling.Lemma))
		t.Logf("unidecode :%s", d.XTokens(ling.Unidecode))

	}
}
