package goenum_test

import (
	"fmt"

	"github.com/oinume/goenum"
)

type LangsEnum struct {
	Go     int
	Python int
	Ruby   int
	Haskel int
}

var (
	lang LangsEnum = LangsEnum{
		Go:     1,
		Python: 2,
		Ruby:   3,
		Haskel: 4,
	}
	langs = goenum.EnumerateStruct(&lang)
)

// Enumerate names
func ExampleEnum_Names() {
	fmt.Printf("%v\n", langs.Names())
	// Output: [Go Python Ruby Haskel]
}

// Enumerate values
func ExampleEnum_Values() {
	fmt.Printf("%v\n", langs.Values())
	// Output: [1 2 3 4]
}

// Obtain name from value
func ExampleEnum_MustName() {
	fmt.Println(langs.MustName(1))
	// Output: Go
}

// Obtain value from name
func ExampleEnum_MustValue() {
	fmt.Println(langs.MustValue("Python"))
	// Output: 2
}
