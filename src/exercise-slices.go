package main

import (
    "code.google.com/p/go-tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
    y := make([][]uint8, dy)
    for i, _ := range y {
        x := make([]uint8, dx)
        for j, _ := range x {
            x[j] = uint8(i % (j+1))
        }
        y[i] = x
    }
    return y
}

func main() {
    pic.Show(Pic)
}
