package unsafe

import (
	"testing"
)

func TestWrapUnwrapIntegers(t *testing.T) {

	g := Ints([]int{1, 3, 3, 7})

	e := (&g).Wrap()

	(&g).Unwrap(e)

	print()

	if g[0] != 1 || g[1] != 3 || g[2] != 3 || g[3] != 7 {
		t.Fatal("Nope.")
	}
}

func TestWrapUnwrapInterfaces(t *testing.T) {
	var foo, bar interface{}

	foo = 1337

	big := Stuff(&foo)
	bar = *Unpack(&big)

	if bar.(int) != 1337 {
		t.Fatal("Nope.")
	}
}
