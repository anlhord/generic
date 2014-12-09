// Package reflect provides unsafe wrapping of an arbitrary slice by reference.
package reflect

import (
	"example.com/repo.git/generic"
	"unsafe"
	"reflect"
)

type ptr uintptr

type Stuffer interface {
	Stuff(interface{}) generic.Value
	Unstuff(interface{}, generic.Value)
}

// FIXME if gc eats our data, we may want to stretch the
// byte slice size a bit :) to accomodate the larger size of Ints or other
// types :)

func (i ptr) wrap() (o generic.Value) {
	p := (*reflect.SliceHeader)(unsafe.Pointer(i))
	q := (*reflect.SliceHeader)(unsafe.Pointer(&o))

	*q = *p

	return
}

func (i ptr) unwrap(o generic.Value) {
	p := (*reflect.SliceHeader)(unsafe.Pointer(i))
	q := (*reflect.SliceHeader)(unsafe.Pointer(&o))

	*p = *q
}


func Stuff(src interface{}) generic.Value {
	return Get(src).wrap()
}

func Unstuff(dst interface{}, val generic.Value) {
	Get(dst).unwrap(val)
}

func Get(i interface{}) ptr {
	return ptr(reflect.ValueOf(i).Pointer())
}
