package main

import "fmt"

type Poem struct {
	content []byte
	storage PoemStorage
}

type PoemStorage interface {
	Type() string
	Load(string) []byte
	Save(string, []byte)
}

func NewPoem(ps PoemStorage) *Poem {
	return &Poem{
		content: []byte("I am a poem from a " + ps.Type() + "."),
	}
}

func (p *Poem) Save(name string) {
	p.storage.Save(name, p.content)
}

func (p *Poem) Load(name string) {
	p.content = p.storage.Load(name)
}

func (p *Poem) String() string {
	return string(p.content)
}

func main() {
	fmt.Println("hello world")
}