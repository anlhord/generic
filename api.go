// Package generic provides the generic Type, Value and Is marker.
package generic

// The generic value.
type Value []byte

// The generic Type is a type of a generic value.
type Type [0]Value

// The Is marker references a concrete type. Marks functions and structures.
type Is uintptr

// The The type references the generic Type type.
type The *Type

// The Iface type references the generic interface{} type.
type Iface *interface{}
