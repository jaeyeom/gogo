package ref

import (
	"errors"
	"reflect"
)

func FieldNames(s interface{}) ([]string, error) {
	t := reflect.TypeOf(s)
	if t.Kind() != reflect.Struct {
		return nil, errors.New("FieldNames: s is not a struct")
	}
	names := []string{}
	n := t.NumField()
	for i := 0; i < n; i++ {
		names = append(names, t.Field(i).Name)
	}
	return names, nil
}

func AppendNilError(f interface{}, err error) (interface{}, error) {
	t := reflect.TypeOf(f)
	if t.Kind() != reflect.Func {
		return nil, errors.New("AppendNilError: f is not a function")
	}
	in, out := []reflect.Type{}, []reflect.Type{}
	for i := 0; i < t.NumIn(); i++ {
		in = append(in, t.In(i))
	}
	for i := 0; i < t.NumOut(); i++ {
		out = append(out, t.Out(i))
	}
	out = append(out, reflect.TypeOf((*error)(nil)).Elem())
	funcType := reflect.FuncOf(in, out, t.IsVariadic())
	v := reflect.ValueOf(f)
	funcValue := reflect.MakeFunc(funcType, func(args []reflect.Value) []reflect.Value {
		results := v.Call(args)
		results = append(results, reflect.ValueOf(&err).Elem())
		return results
	})
	return funcValue.Interface(), nil
}
