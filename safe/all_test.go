package safe

import (
	"testing"
)

func TestWrapUnwrapBytes(t *testing.T) {

	g := Bytes([]byte("Gopher"))

	e := (&g).Wrap()

	(&g).Unwrap(e)

	if string(g) != "Gopher" {
		t.Fatal("Nope.")
	}
}
