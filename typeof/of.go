// Package type provides a type Of method for run/compile time type comparison.
package typeof

import (
	"github.com/anlhord/generic"
	"reflect"
)

// Typeof gets the generic.Type from a reference to a generic.Collection
// Reference to a generic.Is is also supported
// generic.Collection must be at least uint16 aligned for a bit stuffing purpose
func Typeof(i interface{}) (t generic.Type) {
	var is *generic.Is
	var col *generic.Collection
	var iface *generic.Interface

	if reflect.TypeOf(i) == reflect.TypeOf(is) {
		return generic.The
	}
	if reflect.TypeOf(i) == reflect.TypeOf(iface) {
		return generic.Iface
	}
	t = generic.Type(reflect.ValueOf(i).Pointer())

	if reflect.TypeOf(i) == reflect.TypeOf(col) {
		if t & generic.Mask == 1 {
			panic("Unaligned short Collection.")
		}
		return t | generic.Col
	}
	return t 
}

// Is compares a two generic Types
func Is(a,b generic.Type) bool {
	if a < 4 || b < 4 {
		return a == b
	}
	return (generic.Mask & a) == (generic.Mask & b)
}
