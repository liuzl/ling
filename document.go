package ling

func NewDocument(text string) *Document {
	return &Document{Text: text}
}

func (d *Document) String() string {
	return d.Text
}

func (d *Document) NewSpan(start, end int) *Span {
	tokenCnt := len(d.Tokens)
	if tokenCnt == 0 {
		return nil
	}
	if start < 0 {
		start = 0
	}
	if end > tokenCnt {
		end = tokenCnt
	}
	return &Span{Doc: d, Start: start, End: end}
}

func (d *Document) XTokens(anno string) []string {
	var ret []string
	for _, token := range d.Tokens {
		t, has := token.Annotations[anno]
		if !has {
			t = token.Annotations["norm"]
		}
		ret = append(ret, t)
	}
	return ret
}

func (d *Document) XRealTokens(anno string) []string {
	var ret []string
	for _, token := range d.Tokens {
		if token.Type == Space {
			continue
		}
		t, has := token.Annotations[anno]
		if !has {
			t = token.Annotations["norm"]
		}
		ret = append(ret, t)
	}
	return ret
}
