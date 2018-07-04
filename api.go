package ling

import (
	"fmt"
)

type Pipeline struct {
	Annotators []string

	taggers []Processor
}

func (p *Pipeline) AddTagger(t Processor) error {
	if t == nil {
		return fmt.Errorf("cannot add nil tagger!")
	}
	p.taggers = append(p.taggers, t)
	return nil
}

func (p *Pipeline) Annotate(d *Document) error {
	err := Processors["_"].Process(d)
	if err != nil {
		return err
	}
	for _, anno := range p.Annotators {
		if err = Processors[anno].Process(d); err != nil {
			return err
		}
	}

	for _, tagger := range p.taggers {
		if err = tagger.Process(d); err != nil {
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
	return &Pipeline{Annotators: annotators}, nil
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
