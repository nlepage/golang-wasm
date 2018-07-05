package bind

import (
	"reflect"
	"syscall/js"
)

func isGetter(t string, ft reflect.Type) bool {
	return isProperty(t) && ft.NumIn() == 0 && ft.NumOut() == 1
}

func bindGetter(name string, t reflect.Type, parent func() js.Value) reflect.Value {
	switch t.Kind() {
	case reflect.Float64:
		return reflect.ValueOf(func() float64 { return parent().Get(name).Float() })
		//TODO manage Float32
	case reflect.Int:
		return reflect.ValueOf(func() int { return parent().Get(name).Int() })
		//TODO manage other sizes of ints
	case reflect.Bool:
		return reflect.ValueOf(func() bool { return parent().Get(name).Bool() })
	case reflect.String:
		return reflect.ValueOf(func() string { return parent().Get(name).String() })
	case reflect.Struct:
		ft := reflect.FuncOf([]reflect.Type{}, []reflect.Type{t}, false)
		return reflect.MakeFunc(ft, func(args []reflect.Value) []reflect.Value {
			//TODO use sync.Once and keep the value ?
			v := reflect.New(t)
			Bind(v.Interface(), func() js.Value { return parent().Get(name) })
			return []reflect.Value{reflect.ValueOf(v.Elem().Interface())}
		})
	}

	panic("FIXME") //FIXME
}
