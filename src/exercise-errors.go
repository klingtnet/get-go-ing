package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// using `e` in Sprintf would cause an infinte loop, because
	// the `fmt` package will look for the error interface when printing values
	// and therefore calls the Error method in an infinite loop
	x := float64(e)
	return fmt.Sprintf("Can't calculate the square root of `%v`, use complex numbers!", x)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
