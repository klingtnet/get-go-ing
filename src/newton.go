package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    const Δ = 0.001
    z   := 1.
    zz  := z*2
    for math.Abs(zz-z) > Δ {
        zz = z
        z = newt(x, z)
    }
    return z
}

func newt(x float64, z float64) float64 {
    return z - (z*z - x)/2*z
}

func main() {
    fmt.Println(Sqrt(2))
}