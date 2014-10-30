package goenum

import (
	"fmt"
	"reflect"
)

type Enum struct {
	structValue        interface{}
	reflectStructValue reflect.Value  // TODO: Not used yet. For performance
	mapValue           map[int]string // TODO: Implement map to Enum
}

type Enumerator interface {
	Enum() Enum
}

func EnumerateStruct(value interface{}) Enum {
	return Enum{
		structValue:        value,
		reflectStructValue: reflect.Indirect(reflect.ValueOf(value)),
	}
}

func (e Enum) Names() []string {
	if e.structValue != nil {
		value := reflect.Indirect(reflect.ValueOf(e.structValue))
		names := make([]string, value.Type().NumField())
		for i := 0; i < value.Type().NumField(); i++ {
			names[i] = value.Type().Field(i).Name
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
		value := reflect.Indirect(reflect.ValueOf(e.structValue))
		values := make([]int, value.Type().NumField())
		for i := 0; i < value.Type().NumField(); i++ {
			values[i] = int(value.Field(i).Int())
		}
		return values
	} else {
		panic("Not implemented yet")
	}
}

func (e Enum) NameValues() map[int]string {
	nameValues := make(map[int]string)
	value := reflect.Indirect(reflect.ValueOf(e.structValue))
	for i := 0; i < value.Type().NumField(); i++ {
		nameValues[int(value.Field(i).Int())] = value.Type().Field(i).Name
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
