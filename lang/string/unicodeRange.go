package main

import "fmt"

func main() {

	s := "abc"

	for i := 0; i < len(s); i++ {
		fmt.Print(s[i], " ")
	}
	fmt.Println()

	for i := 0; i < len(s); i++ {
		fmt.Print(string(s[i]), " ")
	}
	fmt.Println()
	// OUTPUT:
	// 97 98 99
	// a b c

	u := "\u1EA0\u1EA2\u1EA4"

	for i := 0; i < len(u); i++ {
		fmt.Print(u[i], " ")
	}
	fmt.Println()

	for i := 0; i < len(u); i++ {
		fmt.Print(string(u[i]), " ")
	}
	fmt.Println()
	// OUTPUT:
	// 225 186 160 225 186 161 225 186 162
	// á º   á º ¡ á º ¢

	for _, cp := range u {
		fmt.Print(cp, " ")
	}
	fmt.Println()

	for _, cp := range u {
		fmt.Print(string(cp), " ")
	}
	fmt.Println()
	// OUTPUT:
	// 7840 7842 7844
	// Ạ Ả Ấ

}
