package main

import (
    "fmt"
    "math/rand"
    "math"
    "time"
)

func main() {
    rand.Seed(time.Now().UTC().UnixNano())
    fmt.Printf("My favorite number is %d and not %g!\n", rand.Intn(1000), math.Pi)
}
