// Package low are low level methods to cast anything to and from generic.Value
package low

import (
	"github.com/anlhord/generic"
	"unsafe"
	"reflect"
)

// T transforms object to a **generic.Value
func T(slice interface{}) **generic.Value {
	return (**generic.Value)(unsafe.Pointer(reflect.ValueOf(slice).Pointer()))
}

// Y transforms interface to a *generic.Value
func Y(iface interface{}) *generic.Value {
	var tmp generic.Value
	*(*interface{})(unsafe.Pointer(&tmp)) = *(I(iface))
	return &tmp
}
// I turns object reference into interface{}
func I(iface interface{}) *interface{} {
	return (*interface{})(unsafe.Pointer(reflect.ValueOf(iface).Pointer()))
}

// S checks if the input is a pointer to slice
func S(gvalue interface{}) bool {
	var i *interface{}
	return reflect.TypeOf(gvalue) != reflect.TypeOf(i)
}
