package main

import "fmt"

func main() {
	// notebook := NewNotebook()
	napkin := NewNapkin()

	poem := NewPoem(napkin)
	poem.Save("The first Poem")

	poem = NewPoem(napkin)
	poem.Load("The first Poem")
	fmt.Println(poem)
}

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
		storage: ps,
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

// The Notebook is the storage device going to be injected on the Poem.
// Notebook can store many poems (key, value)
type Notebook struct {
	poems map[string][]byte
}

func NewNotebook() *Notebook {
	return &Notebook{
		poems: map[string][]byte{},
	}
}

func (n *Notebook) Save(name string, contents []byte) {
	n.poems[name] = contents
}

func (n *Notebook) Load(name string) []byte {
	return n.poems[name]
}

// Type returns the description of the storage type.
func (n *Notebook) Type() string {
	return "Notebook"
}

// Napkin is an storage device for a single poem
type Napkin struct {
	poem []byte
}

func NewNapkin() *Napkin {
	return &Napkin{
		poem: []byte{},
	}
}

func (n *Napkin) Save(name string, contents []byte) {
	n.poem = contents
}

func (n *Napkin) Load(name string) []byte {
	return n.poem
}

func (n *Napkin) Type() string {
	return "Napkin"
}