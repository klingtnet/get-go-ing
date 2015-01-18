# Get Go'ing

## What?

Playground for my [first steps in the golang](https://tour.golang.org/list).

## About Go

- Go programs are made out of **packages**
- the `main` method must be in the **main** package
- inside of the **import** statement are the packages specified that should be imported
    - the last element of the *import path* is the package name, by convention. `math/rand` imports the files from `math` that begin with `package rand`
- in Go, a name is exported if it begins with a capital letter, f.e. `math.Pi`

### Functions

- function definitions start with `func` followed by the function name, the parameter list and the return value
    - as opposed to C, the parameter name comes before the type, f.e. `x int`
    - [here is](golang.org/doc/articles/gos_declaration_syntax.html) why they choosed this syntax
    - if two or more consecutive parameters share the same type, you can omit the it from all but the last
    - a function can return *any* number of values (like tuples in python)
- **strings** are enquoted by doublequotes `"`

```go
func add(a, b int) int {
    return a+b
}
```

### Variables

- the **var** statement declares a list of variables with the type last
    - it is allowed on *function* and *package* level (global)
    - examples:
        - `var a, b bool`
        - initalizers can be used like this:
            - `var a, b, c = true, false, "hej!"`
            - note that the type can be omitted if the initializer is present
            - each variable from the initializer list can have a different type
- inside a function the **short assignment** statement can be used: `a := 100`


## Tips

- to build & run a Go file in one step use `go run file.go`
- Go files can be formatted automatically using the `gofmt` tool. On default the formatted code is written to `stdout`, to overwrite the source file use `gofmt -w file.go`.
- the execution environment of a compiled program is deterministic, thus a *random generator* for example has to be seeded, otherwise it will deliver the same number on every run of the program