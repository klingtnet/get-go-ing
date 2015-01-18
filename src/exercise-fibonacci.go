package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var f0, f1 = 0, 1
	return func() int {
		f1, f0 = f0, f1
		f1 = f1 + f0
		return f0 + f1
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
