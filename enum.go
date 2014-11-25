package goenum

import (
	"fmt"
	"reflect"
)

type enumElement struct {
	value int
	name  string
	alias string
}

// struct value of Enum
type Enum struct {
	structValue        interface{}
	reflectStructValue reflect.Value  // TODO: Not used yet. For performance
	mapValue           map[int]string // TODO: Implement map to Enum
	valuesSlice        []int
	elementsMap        map[int]enumElement
}

type Enumerator interface {
	Enum() Enum
}

// Convert an argument 'target' to goenum.Enum
func EnumerateStruct(target interface{}) Enum {
	if target == nil {
		panic("Argument 'target' is nil")
	}

	enum := Enum{
		structValue:        target,
		reflectStructValue: reflect.Indirect(reflect.ValueOf(target)),
	}

	value := enum.reflectStructValue
	enum.valuesSlice = make([]int, value.Type().NumField())
	enum.elementsMap = make(map[int]enumElement, value.Type().NumField())
	for i := 0; i < value.Type().NumField(); i++ {
		key := int(value.Field(i).Int())
		fieldType := value.Type().Field(i)
		enum.valuesSlice[i] = key
		enum.elementsMap[key] = enumElement{
			value: key,
			name:  fieldType.Name,
			alias: fieldType.Tag.Get("goenum"),
		}
	}

	return enum
}

// Return enum names
func (e Enum) Names() []string {
	if e.structValue != nil {
		names := make([]string, len(e.valuesSlice))
		for i := 0; i < len(names); i++ {
			names[i] = e.elementsMap[e.valuesSlice[i]].name
		}
		return names
	} else if e.mapValue != nil {
		panic("Not implemented yet")
	} else {
		panic("Does'nt reach here")
	}
}

// Return enum values
func (e Enum) Values() []int {
	if e.structValue != nil {
		values := make([]int, len(e.valuesSlice))
		copy(values, e.valuesSlice)
		return values
	} else {
		panic("Not implemented yet")
	}
}

// Return enum aliases
func (e Enum) Aliases() []string {
	if e.structValue != nil {
		aliases := make([]string, len(e.valuesSlice))
		for i := 0; i < len(aliases); i++ {
			aliases[i] = e.elementsMap[e.valuesSlice[i]].alias
		}
		return aliases
	} else {
		panic("Not implemented yet")
	}
}

// Return enum values/names as map
func (e Enum) NameValues() map[int]string {
	nameValues := make(map[int]string)
	for i := 0; i < len(e.valuesSlice); i++ {
		value := e.valuesSlice[i]
		name := e.elementsMap[value].name
		nameValues[value] = name
	}
	return nameValues
}

// Return the name for given value. Returns empty string and false if not found.
func (e Enum) Name(value int) (string, bool) {
	nameValues := e.NameValues()
	if v, ok := nameValues[value]; ok {
		return v, true
	} else {
		return "", false
	}
}

// Return the name for given value. panic() if not found.
func (e Enum) MustName(value int) string {
	name, has := e.Name(value)
	if !has {
		panic(fmt.Sprintf("No name for %d", value))
	}
	return name
}

// Return the value for given name. Returns -1 and false if not found.
func (e Enum) Value(name string) (int, bool) {
	for v, n := range e.NameValues() {
		if n == name {
			return v, true
		}
	}
	return -1, false
}

// Return the value for given name. panic() if not found.
func (e Enum) MustValue(name string) int {
	value, has := e.Value(name)
	if !has {
		panic(fmt.Sprintf("No value for %s", name))
	}
	return value
}

// Return the value for given alias. Returns -1 and false if not found.
func (e Enum) ValueForAlias(alias string) (int, bool) {
	for value, element := range e.elementsMap {
		if element.alias == alias {
			return value, true
		}
	}
	return -1, false
}

// Return the value for given alias. panic() if not found.
func (e Enum) MustValueForAlias(alias string) int {
	value, has := e.ValueForAlias(alias)
	if !has {
		panic(fmt.Sprintf("No value for %s", alias))
	}
	return value
}

// Return the alias for given value. panic() if not found.
func (e Enum) Alias(value int) (string, bool) {
	if element, ok := e.elementsMap[value]; ok {
		return element.alias, true
	} else {
		return "", false
	}
}

func (e Enum) MustAlias(value int) string {
	alias, has := e.Alias(value)
	if !has {
		panic(fmt.Sprintf("No alias for %d", value))
	}
	return alias
}
