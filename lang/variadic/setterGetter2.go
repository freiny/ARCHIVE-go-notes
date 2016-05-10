package main

import "fmt"

func main() {

	p := Person{name: "pat"}

	fmt.Println(p.Name(), p.Age())
	// OUTPUT: carl 0

	fmt.Println(p.Name("joe"), p.Age(21))
	// OUTPUT: joe 21

	fmt.Println(p.Name("jan"), p.Age(22))
	// OUTPUT: jan 22

	fmt.Println(p.Name(), p.Age())
	// OUTPUT: jan 22

}

type Person struct {
	name string
	age  int
}

func (p *Person) Name(s ...string) string {
	if s != nil {
		p.name = s[0]
	}
	return p.name
}

func (p *Person) Age(age ...int) int {
	if age != nil {
		p.age = age[0]
	}
	return p.age
}
