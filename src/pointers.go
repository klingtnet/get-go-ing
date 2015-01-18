package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	var p *int
	i := 42
	p = &i
	fmt.Println(*p)
	sp := &Vertex{X: 4, Y:3}
	// indirection/dereferencing trough pointer is transparent
	fmt.Printf("%T %d, %d\n", sp, sp.X, sp.Y)
}
