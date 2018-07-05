package bind

import (
	"reflect"
	"strings"
	"syscall/js"
)

const functionSuffix = "()"

var mapReturns = map[reflect.Kind]func(js.Value) []reflect.Value{
	reflect.Float64: returnFloat,
	reflect.Int:     returnInt,
	reflect.Bool:    returnBool,
	reflect.String:  returnString,
}

type function struct {
	fn        func() func(...interface{}) js.Value
	mapReturn func(js.Value) []reflect.Value
}

func isFunction(t string) bool {
	return strings.HasSuffix(t, functionSuffix)
}

func bindFunction(tag string, t reflect.Type, parent func() js.Value) reflect.Value {
	return reflect.MakeFunc(t, newFunction(tag, t, parent).call)
}

func newFunction(tag string, t reflect.Type, parent func() js.Value) function {
	name := strings.TrimSuffix(tag, functionSuffix)
	fn := func() func(...interface{}) js.Value { return parent().Get(name).Invoke }

	//FIXME check func return type

	var mapReturn func(js.Value) []reflect.Value

	if t.NumOut() == 0 {
		mapReturn = returnVoid
	} else {
		var ok bool
		mapReturn, ok = mapReturns[t.Out(0).Kind()]
		if !ok {
			panic("FIXME") //FIXME
		}
	}

	return function{fn, mapReturn}
}

func (f function) call(argValues []reflect.Value) []reflect.Value {
	args := make([]interface{}, len(argValues))
	for i, argValue := range argValues {
		//TODO if argument is of kind struct, map it to a new JS object...

		args[i] = argValue.Interface()
	}

	return f.mapReturn(f.fn()(args...))
}

func returnVoid(_ js.Value) []reflect.Value {
	return []reflect.Value{}
}

func returnFloat(v js.Value) []reflect.Value {
	return []reflect.Value{reflect.ValueOf(v.Float())}
}

func returnInt(v js.Value) []reflect.Value {
	return []reflect.Value{reflect.ValueOf(v.Int())}
}

func returnBool(v js.Value) []reflect.Value {
	return []reflect.Value{reflect.ValueOf(v.Bool())}
}

func returnString(v js.Value) []reflect.Value {
	return []reflect.Value{reflect.ValueOf(v.String())}
}
