// Package safe provides the wrapping of byte slice.
package safe

import "github.com/anlhord/generic"

type Wrapper interface {
	Wrap() generic.Value
	Unwrap(generic.Value)
}

type Packer interface {
	Pack() *generic.Value // like wrap, may be faster
	Unpack(o *generic.Value) *interface{}
}

type Stuffer interface {
	Stuff(i *interface{}) generic.Value
}

func (b *Bytes) Wrap() generic.Value {
	return generic.Value(*b)
}

func (b *Bytes) Unwrap(o generic.Value) {
	*b = Bytes(o)
}
