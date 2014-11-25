package goenum

import (
	"fmt"
	"reflect"
	"testing"
)

type LangsEnum struct {
	Go     int `goenum:"go"`
	Python int `goenum:"python"`
	Ruby   int `goenum:"ruby"`
	Java   int `goenum:"java"`
}

var langs LangsEnum = LangsEnum{
	Go:     1,
	Python: 2,
	Ruby:   3,
	Java:   4,
}

func (e LangsEnum) Enum() Enum {
	// Convert 'LangsEnum' to 'Enum'
	return EnumerateStruct(&e)
	// TODO: pass map[int]string
	// TODO: want to parse AST of constants and generate code.
}

func TestNames(t *testing.T) {
	expected := []string{"Go", "Python", "Ruby", "Java"}
	actual := langs.Enum().Names()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\nexpect %v\nactual %v", actual, expected)
	}
	//fmt.Println(langs.Enum().Names())
}

func TestValues(t *testing.T) {
	expected := []int{1, 2, 3, 4}
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

	_, has := langs.Enum().Name(0)
	if has {
		t.Errorf("Must be has = false")
	}

	//fmt.Println(langs.Enum().Name(1))
}

func TestMustName(t *testing.T) {
	name := langs.Enum().MustName(langs.Go)
	if name != "Go" {
		t.Errorf("\nexpect %v\nactual %v", "Go", name)
	}
}

func TestValue(t *testing.T) {
	enum := langs.Enum()
	value, _ := enum.Value("Python")
	if value != langs.Python {
		t.Errorf("\nexpect %v\nactual %v", langs.Python, value)
	}

	value, has := enum.Value("ObjectiveC")
	if has {
		t.Errorf("Must be has = false")
	}
	if value != -1 {
		t.Errorf("value must be -1 when not found. value = %d", value)
	}
}

func TestMustValue(t *testing.T) {
	value := langs.Enum().MustValue("Python")
	if value != langs.Python {
		t.Errorf("\nexpect %v\nactual %v", langs.Python, value)
	}
}

type AliasTest struct {
	Exists    int `goenum:"exists"`
	NotExists int
}

func (a AliasTest) Enum() Enum {
	return EnumerateStruct(&a)
}

var aliasTest AliasTest = AliasTest{
	Exists:    1,
	NotExists: 2,
}

func TestAlias(t *testing.T) {
	enum := langs.Enum()
	alias, _ := enum.Alias(langs.Ruby)
	if alias != "ruby" {
		t.Errorf("\nexpect %v\nactual %v", "ruby", alias)
	}

	_, has := enum.Alias(0)
	if has {
		t.Errorf("Must be has = false")
	}

	//fmt.Println(langs.Enum().Name(1))

	e := aliasTest.Enum()
	e.Alias(2)
	if alias2, has := e.Alias(2); has && alias2 != "" {
		t.Errorf("\nexpect %v\nactual %v", "", alias2)
	}
}

func TestMustAlias(t *testing.T) {
	alias := langs.Enum().MustAlias(3)
	if alias != "ruby" {
		t.Errorf("\nexpect %v\nactual %v", "ruby", alias)
	}
}

func TestValueForAlias(t *testing.T) {
	enum := langs.Enum()
	value, _ := enum.ValueForAlias("java")
	if value != langs.Java {
		t.Errorf("\nexpect %v\nactual %v", langs.Java, value)
	}

	_, has := enum.ValueForAlias("python3")
	if has {
		t.Errorf("Must be has = false")
	}
}

func TestMustValueForAlias(t *testing.T) {
	value := langs.Enum().MustValueForAlias("java")
	if value != langs.Java {
		t.Errorf("\nexpect %v\nactual %v", langs.Java, value)
	}
}

func TestAliases(t *testing.T) {
	expected := []string{"go", "python", "ruby", "java"}
	actual := langs.Enum().Aliases()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\nexpect %v\nactual %v", actual, expected)
	}
}

func TestAliasesExists(t *testing.T) {
	expected := []string{"exists", ""}
	actual := aliasTest.Enum().Aliases()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\nexpect %v\nactual %v", actual, expected)
	}
}
