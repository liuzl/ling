package ling

import (
	"encoding/json"
	"fmt"
	"net/url"

	"crawler.club/dl"
	"github.com/juju/errors"
	strutils "github.com/torden/go-strutil"
)

var strvalidator = strutils.NewStringValidator()

// ApiTagger via http interface
type ApiTagger struct {
	apiAddr string
}

// Entity stores the NER entity
type Entity struct {
	Text  string      `json:"text"`
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
	Start int         `json:"start"`
	End   int         `json:"end"`
}

// NewApiTagger returns a new tagger
func NewApiTagger(addr string) (*ApiTagger, error) {
	if !strvalidator.IsValidURL(addr) {
		return nil, fmt.Errorf("invalid url: %s", addr)
	}
	return &ApiTagger{apiAddr: addr}, nil
}

// Process the input document
func (t *ApiTagger) Process(d *Document) error {
	if d == nil || len(d.Text) == 0 {
		return nil
	}
	if len(d.Tokens) == 0 {
		return fmt.Errorf("tokenization required")
	}
	v := url.Values{}
	v.Add("text", d.Text)
	req := &dl.HttpRequest{
		Url:      t.apiAddr,
		Method:   "POST",
		PostData: v.Encode(),
	}
	res := dl.Download(req)
	if res.Error != nil {
		return errors.Annotate(res.Error, "ApiTagger dl.Download")
	}
	var entities []Entity
	if err := json.Unmarshal(res.Content, &entities); err != nil {
		return errors.Annotate(err, "ApiTagger json.Unmarshal")
	}
	for _, entity := range entities {
		if len(entity.Text) < 1 {
			continue
		}
		start := -1
		end := -1
		for _, token := range d.Tokens {
			if token.StartByte == entity.Start {
				start = token.I
			}
			if token.EndByte == entity.End {
				end = token.I + 1
			}
		}
		if start == -1 || end == -1 {
			continue
		}
		span := &Span{Doc: d, Start: start, End: end,
			Annotations: map[string]interface{}{
				"from":  "api",
				"value": map[string]interface{}{entity.Type: entity.Value},
			},
		}
		d.Spans = append(d.Spans, span)
	}
	return nil
}
