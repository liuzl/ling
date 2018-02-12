package ling

type Pipe interface {
	Process(d *Document) error
}
