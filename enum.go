package goenum

import (
	"fmt"
	"reflect"
)

type enumElement struct {
	value   int
	name    string
	aliases []string
}

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
		enum.valuesSlice[i] = key
		enum.elementsMap[key] = enumElement{
			value: key,
			name:  value.Type().Field(i).Name,
			// TODO: aliases
		}
	}

	return enum
}

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

func (e Enum) Values() []int {
	if e.structValue != nil {
		values := make([]int, len(e.valuesSlice))
		copy(values, e.valuesSlice)
		return values
	} else {
		panic("Not implemented yet")
	}
}

func (e Enum) NameValues() map[int]string {
	nameValues := make(map[int]string)
	for i := 0; i < len(e.valuesSlice); i++ {
		value := e.valuesSlice[i]
		name := e.elementsMap[value].name
		nameValues[value] = name
	}
	return nameValues
}

func (e Enum) Name(value int) (string, bool) {
	nameValues := e.NameValues()
	if v, ok := nameValues[value]; ok {
		return v, true
	} else {
		return "", false
	}
}

func (e Enum) MustName(value int) string {
	name, has := e.Name(value)
	if !has {
		panic(fmt.Sprintf("No name for %d", value))
	}
	return name
}

func (e Enum) Value(name string) (int, bool) {
	for v, n := range e.NameValues() {
		if n == name {
			return v, true
		}
	}
	return -1, false
}

func (e Enum) MustValue(name string) int {
	value, has := e.Value(name)
	if !has {
		panic(fmt.Sprintf("No value for %s", name))
	}
	return value
}
