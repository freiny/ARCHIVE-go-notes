package main

import "fmt"

func main() {

	speakers := []Speaker{
		Dog{"Fido"},
		Cat{"Milo"},
	}

	for _, v := range speakers {
		SpeakToUs(v)
	}
	// OUTPUT:
	// Fido: Ruff!
	// Milo: Meow!
}

type Dog struct {
	name string
}

func (d Dog) GetName() string {
	return d.name
}

func (d Dog) Speak() string {
	return "Ruff!"
}

type Cat struct {
	name string
}

func (c Cat) GetName() string {
	return c.name
}

func (c Cat) Speak() string {
	return "Meow!"
}

type Speaker interface {
	GetName() string
	Speak() string
}

func SpeakToUs(s Speaker) {
	fmt.Println(s.GetName()+":", s.Speak())
}
