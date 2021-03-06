// Package generic provides the generic Type, Value, Is and Collection tag.
package generic

// The Type references a generic or concrete handle.
type Type uintptr

// The generic value.
type Value []byte

// The generic Is used to tag various structs, functions to be generic.
type Is [0]Value

// Marker for an user-specific type collections.
const CollectionOf = 0

// Used to make a generic collection.
// Don't forget to pad different collections so they don't share a field
type Collection [0]Value

// The generic Interface.
// Don't forget to pad different collections so they don't share a field
type Interface [1]interface{}

// The The type is a magic value that references the generic Is.
const The = Type(uintptr(2))

// The Col type is a magic value that references the generic Collection.
const Col = Type(uintptr(1))

// The Iface type is a magic value that references a generic interface{} Tag.
const Iface = Type(uintptr(3))

// Used to mask out a magic values from the Type reference.
const Mask = Type(^uintptr(1))
