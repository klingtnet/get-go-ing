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

### Variables & Types

- the **var** statement declares a list of variables with the type last
    - it is allowed on *function* and *package* level (global)
    - examples:
        - `var a, b bool`
        - initalizers can be used like this:
            - `var a, b, c = true, false, "hej!"`
            - note that the type can be omitted if the initializer is present
            - each variable from the initializer list can have a different type
    - var statements can be factored into blocks, similar to the import statement, see [basictypes.go](./src/basictypes.go) for an example
    - variables declared without an explicit initial value will be instantiated with their specific [zero value](https://tour.golang.org/basics/12)
- inside a function the **short assignment** statement can be used: `a := 100`
- **type conversions** can be done with `T(..)`, where `T` is the type and inside of the parantheses is the value to convert, f.e. `float64(128)`

### Constants

- declared using the `const` keyword
- **can't** be declared using the short assigment statement `:=`
- constants can be character, string boolean or numeric values
- numeric-constants are [*high-precision* values](https://tour.golang.org/basics/16)

### Loops

- go has only one looping construct, the `foor` loop
- to emulate a `while` loop leave the *pre* and *post* statements empty: `for ; x < y; {}`, you can even omit out the `;`: `for x < y {}`
- omit the loop conditions and you get an infinite loop: `for {}`

```go
for i := 0; i < 10; i++ {
    sum += i
}
```

### Conditions

- C-like but without the parentheses:

```go
if x < y {
    x++
} else {
    y++
}
```

- you can write a pre-statement before the if-statement
- variables declared in this pre-statement are only visible inside the scope of the if statement

```go
if x := 10; x == 10 {
    fmt.Println("It's only an example.")
}
```

- **switch-case** statements break automatically, unless you specfiy a *falltrough* statement (`default` case)
- the *evaluation order* is from *top to bottom*
- a `switch` without condition is the same as `switch true` and can be used for long if-else chains:

```go
switch {
case t.Hour() < 12:
    fmt.Println("Good morning!")
case t.Hour() < 17:
    fmt.Println("Good afternoon.")
default:
    fmt.Println("Good evening.")
}
```

### Miscellanous

- the `defer` statement defers the execution of a function until the surrounding function returns
- deferred function calls are pushed on a stack and are executed in **LIFO** order
- [more on defer](http://blog.golang.org/defer-panic-and-recover)

## Tips

- to build & run a Go file in one step use `go run file.go`
- Go files can be formatted automatically using the `gofmt` tool. On default the formatted code is written to `stdout`, to overwrite the source file use `gofmt -w file.go`.
- the execution environment of a compiled program is deterministic, thus a *random generator* for example has to be seeded, otherwise it will deliver the same number on every run of the program