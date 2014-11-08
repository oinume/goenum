[![Build Status](http://img.shields.io/travis/oinume/goenum.svg?style=flat)](https://travis-ci.org/oinume/goenum)
[![Coverage](http://img.shields.io/codecov/c/github/oinume/goenum.svg?style=flat)](https://codecov.io/github/oinume/goenum)
[![License](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/oinume/goenum/blob/master/LICENSE)

# goenum

Enum for Go

# Example

```go
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/oinume/goenum"
)

type LangsEnum struct {
	Go     int
	Python int
	Ruby   int
	Haskel int
}

var Langs LangsEnum = LangsEnum{1, 2, 3, 4}

func (e LangsEnum) Enum() goenum.Enum {
	// Make Enum from Langs
	return goenum.EnumerateStruct(&Langs)
}

func main() {
	var langs goenum.Enum = Langs.Enum()

	// Enumerate names
	fmt.Println(langs.Names())
	// --> [Go Python Ruby Haskel]

	// Enumerate values
	fmt.Println(langs.Values())
	// --> [1 2 3 4]

	// Obtain name from value
	fmt.Println(langs.MustName(1))
	// --> Go

	// Obtain value from name
	fmt.Println(langs.MustValue("Python"))
	// --> 2

	if len(os.Args) > 1 {
		// $ go run sample/langs.go 2
		value, _ := strconv.Atoi(os.Args[1])
		switch value {
		case Langs.Go:
			fmt.Println("Gopher")
		case Langs.Python:
			fmt.Println("Pythonista")
		case Langs.Ruby:
			fmt.Println("Rubyist")
		case Langs.Haskel:
			fmt.Println("Haskeler")
		}
	}
}
```
