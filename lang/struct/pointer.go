package main

import "fmt"

type User struct {
	name string
	age  int
}

func (u *User) writable() {
	u.name = "Mutated"
	u.age = 987
}

func (u User) readable() {
	u.name = "Mutated"
	u.age = 987
}

func main() {
	user := User{"James Bond", 26}

	fmt.Println(user)
	// OUTPUT:
	// {James Bond 26}

	user.readable()
	fmt.Println(user)
	// OUTPUT:
	// {James Bond 26}

	user.writable()
	fmt.Println(user)
	// OUTPUT:
	// {Mutated 987}

}
