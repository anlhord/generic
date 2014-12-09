package unsafe

// Ints for receivers
type Ints []int

// Internal interface{} representation. Must match the actual one!
type Iface [2]uintptr
