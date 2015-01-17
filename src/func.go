package main

import (
    "fmt"
)

func main() {
    fmt.Println(add(12, 30))
    fmt.Println(swap("Linz", "Andreas"))
}

func add(a, b int) int {
    return a + b
}

// return any number of arguments
func swap(a, b string) (string, string) {
    return b, a
}

// named return values
// use this only in short functions
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}