package main

import (
	"fmt"
)

type Decimal struct {
	X float64
}

func (v Decimal) Double() float64 {
	return 2 * v.X
}

func (v *Decimal) DoublePR() {
	v.X = 2 * v.X
}

func main() {
	v := Decimal{3.14}
	// call-by-value
	fmt.Println(v, v.Double())
	// use the pointer Receiver
	v.DoublePR()
	// the value of v has changed without explicit assignment
	fmt.Println(v)
}
