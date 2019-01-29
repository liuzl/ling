# ling is a golang toolkit for natural language processing
[![GoDoc](https://godoc.org/github.com/liuzl/ling?status.svg)](https://godoc.org/github.com/liuzl/ling)[![Go Report Card](https://goreportcard.com/badge/github.com/liuzl/ling)](https://goreportcard.com/report/github.com/liuzl/ling)

# Implementation references
## Similar NLP tools
* [Stanford CoreNLP](https://stanfordnlp.github.io/CoreNLP/index.html) Java
* [spaCy](https://spacy.io/) Python
* [lingo](https://github.com/chewxy/lingo) Golang
## Multilingual text toknization
* [Unicode Standard Annex #29](http://www.unicode.org/reports/tr29/)
* [blevesearch segment](https://github.com/liuzl/segment)
## Text normalization
* [Text normalization in Go](https://blog.golang.org/normalization)
## Lemmatization
> 词干提取（stemming）和词形还原（lemmatization）

* [Stemming and lemmatization](https://nlp.stanford.edu/IR-book/html/htmledition/stemming-and-lemmatization-1.html)
* [Lemmatization Lists](http://www.lexiconista.com/datasets/lemmatization/)*<sub>Datasets by MBM </sub>*
* [The UniMorph Project](https://unimorph.github.io/)
* 中文繁简转换
  * [gocc](https://github.com/liuzl/gocc) Golang version OpenCC
  * [OpenCC](https://github.com/BYVoid/OpenCC)
  * [Chinese-Character Jian<=>Fan converting library in Go](https://github.com/go-cc/cc-jianfan)
  * [Traditional and Simplified Chinese Conversion in Go](https://github.com/siongui/gojianfan)
  * [Han unification](https://en.wikipedia.org/wiki/Han_unification)
## Tagging
* Regex tagger
  * [commonregex](https://github.com/mingrammer/commonregex), a collection of common regular expressions for Go.
  * [xurls](https://github.com/mvdan/xurls), a Go package of regex for urls.
## Natural language Detection

`getlang` is much slower than `franco`

* [getlang](https://github.com/rylans/getlang)
* [franco](https://github.com/liuzl/franco)
* [test scripts](https://github.com/liuzl/org_name_parser/blob/master/parse/pprof.sh)
  * franco: Duration: 5.12s, 26.93%
  * getlang: Duration: 11.58s, 59.54%
