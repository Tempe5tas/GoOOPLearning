package Extension

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(name string) {
	p.Speak()
	fmt.Println(" ", name)
}

type Dog struct {
	p *Pet
}

//func (d *Dog) Speak() {
//	d.p.Speak()
//}

//func (d *Dog) SpeakTo(name string) {
//	d.p.SpeakTo(name)
//}

func (d *Dog) Speak() {
	fmt.Print("Woof!")
}

func (d *Dog) SpeakTo(name string) {
	d.Speak()
	fmt.Println(" ", name)
}

func TestExtension(t *testing.T) {
	dog := &Dog{}
	dog.SpeakTo("master")
}
