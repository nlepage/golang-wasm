package bind

import (
	"fmt"
	"reflect"
	"syscall/js"
)

// BindGlobals binds JS identifiers from the global scope into a structure
// This is equivalent to Bind(v, js.Global())
func BindGlobals(v interface{}) error {
	return Bind(v, js.Global)
}

// Bind binds JS identifiers from the given JS parent reference
func Bind(v interface{}, parent func() js.Value) error {
	if err := checkType(v); err != nil {
		return err
	}

	// rv is the reflect.Value of the struct (obtained by resolving the pointer to struct)
	rv := reflect.ValueOf(v).Elem()

	for i, f := range getFields(v) {
		// t is the js struct tag of the field, defaulted to the field name
		tag := f.Tag.Get("js")
		if tag == "" {
			tag = f.Name
		}

		// fv is the reflect.Value of the field
		fv := rv.Field(i)
		// ft is the reflect.Type of the field
		ft := fv.Type()

		// Return an error if the type of the field is not a func type
		if k := ft.Kind(); k != reflect.Func {
			return fmt.Errorf("Field of type %s found, func expected", k)
		}

		// value will be the new value of the field
		var value reflect.Value

		switch {
		//TODO allow js.Value ?
		case isGetter(tag, ft):
			value = bindGetter(tag, ft.Out(0), parent)
		case isSetter(tag, ft):
			value = bindSetter(tag, ft.In(0), parent)
		case isFunction(tag):
			value = bindFunction(tag, ft, parent)
		}
		//FIXME default error

		// Set the new value of the field
		fv.Set(value)
	}

	return nil
}

// checkType checks if v is a pointer to struct
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

// getFields returns the list of fields of the struct in a slice of reflect.StructField
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
