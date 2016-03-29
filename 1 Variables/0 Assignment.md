Variable Declaration and Assignment
<pre><code>
package main

import "fmt"

func main() {

	// declare variable named i with integer type
	// set variable to zero value for int type
	var i int
	fmt.Println(i)
	// OUTPUT:
	// 0

	// assign variable a different value
	i = 4
	fmt.Println(i)
	// OUTPUT:
	// 4

	// declare and assign integer j
	// j's type is inferred from it's value on the right
	j := 9
	fmt.Println(j)
	// OUTPUT:
	// 9

	// multiple declaration
	var a, b, c int
	fmt.Println(a, b, c)
	// OUTPUT:
	// 0 0 0

	// multiple assignment
	// set variable to new values
	a, b, c = 1, 2, 3
	fmt.Println(a, b, c)
	// OUTPUT:
	// 1 2 3

	// multiple declaration and assignment
	d, e, f := 4, 5, 6
	fmt.Println(d, e, f)
	// OUTPUT:
	// 4 5 6

	// multiple declaration and assignment
	// using different types
	g, h, i := 7, "seven", 7.0
	fmt.Println(g, h, i)
	// OUTPUT:
	// 7 seven 7

	// var explicitly states your intention to declare a variable
	var msg1 = "Hello"
	fmt.Println(msg1)
	// OUTPUT:
	// Hello

	// := explicitly states your intention to declare a variable
	msg2 := "GoodBye"
	fmt.Println(msg2)
	// OUTPUT:
	// GoodBye

}
</code></pre>

OUTPUT:
<pre><code>
0
4
9
0 0 0
1 2 3
4 5 6
7 seven 7
Hello
GoodBye
</code></pre>
