package bind

import (
	"fmt"
	"reflect"
	"strings"
	"syscall/js"
)

const functionSuffix = "()"

func BindGlobals(v interface{}) error {
	return Bind(v, js.Global())
}

func Bind(v interface{}, parent js.Value) error {
	if err := checkType(v); err != nil {
		return err
	}

	rv := reflect.ValueOf(v).Elem()

	for i, f := range getFields(v) {
		t := f.Tag.Get("js")

		if t == "" {
			continue
		}

		fv := rv.Field(i)
		ft := fv.Type()

		if k := ft.Kind(); k != reflect.Func {
			return fmt.Errorf("Field of type %s found, func expected", k)
		}

		var value reflect.Value

		switch {
		case isGetter(t, ft):
			value = bindGetter(t, ft.Out(0), parent)
		case isSetter(t, ft):
			value = bindSetter(t, ft.In(0), parent)
		case isFunction(t):
			value = bindFunction(strings.TrimSuffix(t, functionSuffix), ft, parent)
		}

		fv.Set(value)
	}

	return nil
}

func checkType(v interface{}) error {
	t := reflect.TypeOf(v)

	if k := t.Kind(); k != reflect.Ptr {
		return fmt.Errorf("%s received, ptr to struct expected", k)
	}

	if k := t.Elem().Kind(); k != reflect.Struct {
		return fmt.Errorf("ptr to %s received, ptr to struct expected", k)
	}

	return nil
}

func getFields(v interface{}) []reflect.StructField {
	t := reflect.TypeOf(v).Elem()
	num := t.NumField()

	fields := make([]reflect.StructField, num)

	for i := 0; i < num; i++ {
		fields[i] = t.Field(i)
	}

	return fields
}

func isProperty(t string) bool {
	return !isFunction(t)
}

func isGetter(t string, ft reflect.Type) bool {
	return isProperty(t) && ft.NumIn() == 0 && ft.NumOut() == 1
}

func bindGetter(name string, t reflect.Type, parent js.Value) reflect.Value {
	value := parent.Get(name)

	var getter func() interface{}
	switch t.Kind() {
	case reflect.Float32:
	case reflect.Float64:
		getter = func() interface{} { return value.Float() }
	case reflect.Int:
		getter = func() interface{} { return value.Int() }
	case reflect.Bool:
		getter = func() interface{} { return value.Bool() }
	case reflect.String:
		getter = func() interface{} { return value.String() }
	}

	return reflect.MakeFunc(reflect.FuncOf([]reflect.Type{}, []reflect.Type{t}, false), func(_ []reflect.Value) []reflect.Value {
		return []reflect.Value{reflect.ValueOf(getter())}
	})
}

func isSetter(t string, ft reflect.Type) bool {
	return isProperty(t) && ft.NumIn() == 1 && ft.NumOut() == 0
}

func bindSetter(name string, t reflect.Type, parent js.Value) reflect.Value {
	setter := func(v interface{}) { parent.Set(name, v) }

	return reflect.MakeFunc(reflect.FuncOf([]reflect.Type{t}, []reflect.Type{}, false), func(args []reflect.Value) []reflect.Value {
		setter(args[0].Interface())
		return []reflect.Value{}
	})
}

func isFunction(t string) bool {
	return t != "" && strings.HasSuffix(t, functionSuffix)
}

func bindFunction(name string, t reflect.Type, parent js.Value) reflect.Value {
	fn := parent.Get(name).Invoke
	return reflect.MakeFunc(t, func(argValues []reflect.Value) []reflect.Value {
		args := make([]interface{}, len(argValues))
		for i, argValue := range argValues {
			args[i] = argValue.Interface()
		}
		fn(args...)
		return []reflect.Value{} //FIXME manage return value and error
	})
}
