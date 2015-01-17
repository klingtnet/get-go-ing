# Get Go'ing

## What?

Playground for my [first steps in the golang](https://tour.golang.org/list).

## About Go

- Go programs are made out of **packages**
- the `main` method must be in the **main** package
- inside of the **import** statement are the packages specified that should be imported
    - the last element of the *import path* is the package name, by convention. `math/rand` imports the files from `math` that begin with `package rand`
- in Go, a name is exported if it begins with a capital letter, f.e. `math.Pi`
- function definitions start with `func` followed by the function name, the parameter list and the return value
    - as opposed to C, the parameter name comes before the type, f.e. `x int`

```go
func add(a int, b int) int {
    return a+b
}
```


## Tips

- to build & run a Go file in one step use `go run file.go`
- Go files can be formatted automatically using the `gofmt` tool. On default the formatted code is written to `stdout`, to overwrite the source file use `gofmt -w file.go`.
- the execution environment of a compiled program is deterministic, thus a *random generator* for example has to be seeded, otherwise it will deliver the same number on every run of the program