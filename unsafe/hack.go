package unsafe

import (
	"github.com/anlhord/generic"
	"unsafe"
)

// FIXME if gc eats our data, we may want to stretch the
// byte slice size a bit :) to accomodate the larger size of Ints or other
// types :)

func (i *Ints) Pack() (o *generic.Value) {
	return (*generic.Value)(unsafe.Pointer(i))
}

func (i *Ints) Wrap() (o generic.Value) {
	q := (*Ints)(unsafe.Pointer(&o))

	*q = *i

	return o
}

func (i *Ints) Unwrap(o generic.Value) {
	q := (*Ints)(unsafe.Pointer(&o))

	*i = *q
}

func (i *Iface) Pack() (o *generic.Value) {
	return (*generic.Value)(unsafe.Pointer(i))
}

func (i *Iface) Wrap() (o generic.Value) {
	q := (*Iface)(unsafe.Pointer(&o))
	*q = *i
	return o
}

func (i *Iface) Unwrap(o generic.Value) {
	q := (*Iface)(unsafe.Pointer(&o))

	*i = *q
}

func pack(i *interface{}) *Iface {
	return (*Iface)(unsafe.Pointer(i))
}

func Stuff(i *interface{}) generic.Value {
	return pack(i).Wrap()
}

func Unpack(o *generic.Value) *interface{} {
	return (*interface{})(unsafe.Pointer(o))
}
