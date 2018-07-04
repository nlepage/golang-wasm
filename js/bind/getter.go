package bind

import (
	"reflect"
	"syscall/js"
)

func isGetter(t string, ft reflect.Type) bool {
	return isProperty(t) && ft.NumIn() == 0 && ft.NumOut() == 1
}

func bindGetter(name string, t reflect.Type, parent js.Value) reflect.Value {
	value := parent.Get(name)

	switch t.Kind() {
	case reflect.Float64:
		return reflect.ValueOf(value.Float)
		//TODO manage Float32
	case reflect.Int:
		return reflect.ValueOf(value.Int)
		//TODO manage other sizes of ints
	case reflect.Bool:
		return reflect.ValueOf(value.Bool)
	case reflect.String:
		return reflect.ValueOf(value.String)
	case reflect.Struct:
		ft := reflect.FuncOf([]reflect.Type{}, []reflect.Type{t}, false)
		return reflect.MakeFunc(ft, func(args []reflect.Value) []reflect.Value {
			//TODO use sync.Once and keep the value ?
			v := reflect.New(t)
			Bind(v.Interface(), value)
			return []reflect.Value{reflect.ValueOf(v.Elem().Interface())}
		})
	}

	panic("FIXME") //FIXME
}
