# 1. Init step
## Instalation
- Install Go
```
https://go.dev/doc/install
```
![Install](https://github.com/dong-tran/goinaction/blob/master/assets/install.png)

## Verify installing
```
go version
```
`go version go1.19.2 darwin/arm64`


# 2. Buildin & Standard library

## Buildin
```
https://pkg.go.dev/builtin@go1.19.2
```
![Buildin](https://github.com/dong-tran/goinaction/blob/master/assets/buildin.png)

## Standard library
```
https://pkg.go.dev/std
```
- compress
- database
- encoding
- errors
- html
- math
- net
- os
- strconv
- strings
- time

# 3. Create module
1. Create folder `goinaction`
```
mkdir goinaction && cd goinaction
```
2. Init module
```
go mod init <module-name>
```
- `go.mod`
- `go.sum`
- `main.go`
```
go get <module-name>
```

# 4. Run a module
- Dev
```
go run main.go
```
- Compile & execute
```
go build -o ~/executable-file
```
```
/bin/bash ~/executable-file
```

# 5. Define `type`
- Interface
```
type Vehicle interface {
	Run()
	Stop()
}
```

- `struct`
```
type Vehicle struct {
	Weight int
    Speed  int
}
```

- simple type
```
type MyNumber int
```
- pointer
```
type MyNumber *int
```
- function
```
package main

import "fmt"

type MyFunc func(a, b int) int

func (f MyFunc) Multiple(a, b, c int) int {
	return f(a, b) * c
}

func main() {
	var a = MyFunc(func(a, b int) int {
		return a + b
	})
	fmt.Printf("a=%d", a.Multiple(1, 3, 5))
    var b = MyFunc(func(a, b int) int {
		return a * b
	})
    fmt.Printf("b=%d", a.Multiple(1, 3, 5))
}

```
https://go.dev/play/

- Method
```
type Vehicle struct {
	Weight int
    Speed  int
}

func (v Vehicle) Run() string {
    return fmt.Printf("Running... speed=%d", v.Speed)
}
```

# 6. Not included
- Source Code Elements
- Keywords and Identifiers
- Basic Types and Their Value Literals
- Constants and Variables
- Common Operators.
- Function Declarations and Calls
- Code Packages and Package Imports
- Expressions, Statements and Simple Statements
- Basic Control Flows
- Goroutines, Deferred Function Calls and Panic/Recover
- Memory Related
- Compiler

# 7. OOP?
