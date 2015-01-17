package main

import (
    "fmt"
)

func main() {
    fmt.Println(add(12, 30))
}

func add(a int, b int) int {
    return a + b
}