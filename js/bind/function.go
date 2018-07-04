package bind

import (
	"reflect"
	"strings"
	"syscall/js"
)

const functionSuffix = "()"

func isFunction(t string) bool {
	return strings.HasSuffix(t, functionSuffix)
}

// FIXME refactor
func bindFunction(tag string, t reflect.Type, parent js.Value) reflect.Value {
	name := strings.TrimSuffix(tag, functionSuffix)
	fn := parent.Get(name).Invoke

	//FIXME check func return type

	var returnMapper func(js.Value) []reflect.Value

	switch t.NumOut() {
	case 0:
		returnMapper = func(_ js.Value) []reflect.Value { return []reflect.Value{} }
	case 1:
		//TODO allow js.Value ?
		switch t.Out(0).Kind() {
		case reflect.Float64:
			returnMapper = func(v js.Value) []reflect.Value { return []reflect.Value{reflect.ValueOf(v.Float())} }
			//TODO manage Float32
		case reflect.Int:
			returnMapper = func(v js.Value) []reflect.Value { return []reflect.Value{reflect.ValueOf(v.Int())} }
			//TODO manage other sizes of ints
		case reflect.Bool:
			returnMapper = func(v js.Value) []reflect.Value { return []reflect.Value{reflect.ValueOf(v.Bool())} }
		case reflect.String:
			returnMapper = func(v js.Value) []reflect.Value { return []reflect.Value{reflect.ValueOf(v.String())} }
		}
	}

	return reflect.MakeFunc(t, func(argValues []reflect.Value) []reflect.Value {
		args := make([]interface{}, len(argValues))
		for i, argValue := range argValues {
			//TODO if argument is of kind struct, map it to a new JS object...

			args[i] = argValue.Interface()
		}

		return returnMapper(fn(args...))
	})
}
