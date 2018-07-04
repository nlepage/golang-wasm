package bind

import (
	"reflect"
	"syscall/js"
)

func isSetter(t string, ft reflect.Type) bool {
	return isProperty(t) && ft.NumIn() == 1 && ft.NumOut() == 0
}

type setter struct {
	name   string
	parent js.Value
}

func (s setter) set(v interface{}) {
	s.parent.Set(s.name, v)
}

type floatSetter struct {
	setter
}

func (s floatSetter) set(f float64) {
	s.setter.set(f)
}

type intSetter struct {
	setter
}

func (s intSetter) set(f int) {
	s.setter.set(f)
}

type boolSetter struct {
	setter
}

func (s boolSetter) set(f bool) {
	s.setter.set(f)
}

type stringSetter struct {
	setter
}

func (s stringSetter) set(f string) {
	s.setter.set(f)
}

func bindSetter(name string, t reflect.Type, parent js.Value) reflect.Value {
	s := setter{name, parent}

	switch t.Kind() {
	case reflect.Float64:
		return reflect.ValueOf(floatSetter{s}.set)
	case reflect.Int:
		return reflect.ValueOf(intSetter{s}.set)
	case reflect.Bool:
		return reflect.ValueOf(boolSetter{s}.set)
	case reflect.String:
		return reflect.ValueOf(stringSetter{s}.set)
	}

	panic("FIXME") //FIXME
}
