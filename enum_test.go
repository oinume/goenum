package goenum

import (
	"fmt"
	"reflect"
	"testing"
)

type LangsEnum struct {
	Go int
	Python int
	Ruby int
	Java int
}

var langs LangsEnum = LangsEnum{
	Go: 1,
	Python: 2,
	Ruby: 3,
	Java: 4,
}

func (e LangsEnum) Enum() Enum {
	// Convert 'LangsEnum' to 'Enum'
	return EnumerateStruct(&e)
	// TODO: pass map[int]string
	// TODO: want to parse AST of constants and generate code.
}

func TestNames(t *testing.T) {
	expected := []string{ "Go", "Python", "Ruby", "Java" }
	actual := langs.Enum().Names()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\nexpect %v\nactual %v", actual, expected)
	}
	//fmt.Println(langs.Enum().Names())
}

func TestValues(t *testing.T) {
	expected := []int{ 1, 2, 3, 4 }
	actual := langs.Enum().Values()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\nexpect %v\nactual %v", actual, expected)
	}
	//fmt.Println(actual)
}

func TestNameValue(t *testing.T) {
	fmt.Println(langs.Enum().NameValues())
}

func TestName(t *testing.T) {
	name, _ := langs.Enum().Name(1)
	if name != "Go" {
		t.Errorf("\nexpect %v\nactual %v", "Go", name)
	}

	_, err := langs.Enum().Name(0)
	if err == nil {
		t.Errorf("Must be err != nil")
	}

	//fmt.Println(langs.Enum().Name(1))
}