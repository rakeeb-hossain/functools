# functools

[![Go](https://github.com/rakeeb-hossain/functools/actions/workflows/go.yml/badge.svg)](https://github.com/rakeeb-hossain/functools/actions/workflows/go.yml)

functools is a simple Go library that brings you your favourite functional paradigms without sacrificing type-safety using 
`interface{}` or `reflect`

Made possible by Go 1.18 using the newly introduced generics.

## Features
 
- Any
- All
- Count
- Filter
- ForEach
- Map
- Reduce
- Sum

## Installation

`go get -u github.com/rakeeb-hossain/functools`

## Usage

```go
import (
    "github.com/rakeeb-hossain/functools"
    "fmt"
)

type User struct {
	username     string
	hasPortfolio bool
}

var users = []User{
		{"gopher", true},
		{"rakeeb", false},
		{"jack", true}}

func main() {
    // Count users with linked portfolios
    fmt.Printf("num users with linked portfolios: %d", 
        functools.Count(users, func(u User) bool { return u.hasPortfolio }))

    // Print usernames of users with linked portfolios
    functools.ForEach(
        functools.Filter(users, func(u User) bool { return u.hasPortfolio }),
        func(u User) { fmt.Printf("%s has a linked portfolio\n", u.username) })
}
```

## Documentation

https://pkg.go.dev does not yet support Go 1.18 packages that use generics: https://github.com/golang/go/issues/48264

For now, documentation is provided via comments and by running `go doc -all` from the package directory. 

## Contributing

Please see [CONTRIBUTING](CONTRIBUTING.md)
