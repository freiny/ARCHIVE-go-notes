package main

import "fmt"

type Point struct {
	X int
	Y int
}

type Circle struct {
	Center Point
	R      int
}

type CircleAnon struct {
	Point
	R int
}

func main() {

	// Named field Center, requires c.Center.X to refer to sub struct field X
	c := Circle{Point{1, 2}, 3}
	fmt.Println(c.Center.X, c.Center.Y, c.R)
	// OUTPUT: 1 2 3

	// Anonymous field Point, requires only c.X to refer to sub struct field X
	ca := CircleAnon{Point{1, 2}, 3}
	fmt.Println(ca.X, ca.Y, ca.R)
	// OUTPUT: 1 2 3

}
