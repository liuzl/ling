package ling

import (
	"fmt"
)

// A Pipeline contains configured annotators and taggers for nl processing
type Pipeline struct {
	Annotators []string

	taggers []Processor
}

// AddTagger adds a new processor t to Pipeline p
func (p *Pipeline) AddTagger(t Processor) error {
	if t == nil {
		return fmt.Errorf("cannot add nil tagger")
	}
	p.taggers = append(p.taggers, t)
	return nil
}

// Annotate tags the Document by each configured processors and taggers
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

// AnnotatePro tags the Document by each configured processors and taggers
func (p *Pipeline) AnnotatePro(d *Document, taggers ...Processor) error {
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
	for _, tagger := range taggers {
		if err = tagger.Process(d); err != nil {
			return err
		}
	}
	return nil
}

// NLP returns ling handler with the annotators
func NLP(annotators ...string) (*Pipeline, error) {
	for _, anno := range annotators {
		if Processors[anno] == nil {
			return nil, fmt.Errorf("Processor %s doesn't exists", anno)
		}
	}
	return &Pipeline{Annotators: annotators}, nil
}

// DefaultNLP returns ling handler with norm, lemma, unidecode and regex
func DefaultNLP() (*Pipeline, error) {
	return NLP("norm", "lemma", "unidecode", "regex")
}

// MustNLP is like NLP but panics if the annotators are not correct. It
// simplifies safe initialization of global variables holding ling handler
func MustNLP(annotators ...string) *Pipeline {
	p, err := NLP(annotators...)
	if err != nil {
		panic(err)
	}
	return p
}
