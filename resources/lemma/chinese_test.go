package lemma

import (
	"testing"
)

func TestConvert(t *testing.T) {
	cases := []string{
		`乾隆爺的幹爺爺是誰`,
	}
	for _, c := range cases {
		out, err := Convert(c)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%+v\n", out)
	}
}
