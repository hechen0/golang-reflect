package main

import (
	"fmt"
	"reflect"
	"errors"
)

type Test struct {
	Id   int64
	Name string
}

func main() {
	a := []int{1}
	ReflectAppend(&a, 300)
	fmt.Println(a)

	b := []string{"hello"}
	ReflectAppend(&b, "world")
	fmt.Println(b)

	c := []Test{Test{1, "hello"}}
	ReflectAppend(&c, Test{2, "world"})
	fmt.Println(c)
}

func ReflectAppend(slice interface{}, value interface{}) {
	slicePtr := reflect.ValueOf(slice)
	sliceValue := reflect.Indirect(slicePtr)
	elemValue := reflect.ValueOf(value)
	resultValue := reflect.Append(sliceValue, elemValue)
	sliceValue.Set(resultValue)
}

func ReflectCallFunc(f interface{}, args ...interface{}) (out []interface{}, err error) {
	fv := reflect.ValueOf(f)
	if fv.Kind() != reflect.Func {
		return nil, errors.New("expect func type for f")
	}
	ft := fv.Type()
	margs := ft.NumIn()
	inv := make([]reflect.Value, margs)

	for n := 0; n < margs; n++ {
		if n < len(args) {
			inv[n] = reflect.ValueOf(args[n])
		} else {
			inv[n] = reflect.Zero(ft.In(n))
		}
	}
	outv := fv.Call(inv)
	out = make([]interface{}, ft.NumOut())
	for n := 0; n < ft.NumOut(); n++ {
		out[n] = outv[n].Interface()
	}
	return out, nil
}
