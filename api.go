package ling

import (
	"fmt"
)

type Pipeline struct {
	Annotators []string
}

func (p *Pipeline) Annotate(d *Document) error {
	err := Processors["_"].Process(d)
	if err != nil {
		return err
	}
	for _, anno := range p.Annotators {
		err = Processors[anno].Process(d)
		if err != nil {
			return err
		}
	}
	return nil
}

func NLP(annotators ...string) (*Pipeline, error) {
	for _, anno := range annotators {
		if Processors[anno] == nil {
			return nil, fmt.Errorf("Processor %s doesn't exists", anno)
		}
	}
	return &Pipeline{annotators}, nil
}

func DefaultNLP() (*Pipeline, error) {
	return NLP("norm", "lemma", "unidecode", "regex")
}

func MustNLP(annotators ...string) *Pipeline {
	p, err := NLP(annotators...)
	if err != nil {
		panic(err)
	}
	return p
}
