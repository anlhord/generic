package reflect

import (
	"example.com/repo.git/generic"
	"testing"
)

type Wrapper interface {
	Wrap() generic.Value
	Unwrap(generic.Value)
}

type hipster struct {
	age       int
	moustache bool
	fixiebike   bool
	glasses   bool
	gender    byte
}

func TestWrapCustomSlice(t *testing.T) {
	var comrades []hipster
	friends := []hipster{{19, false, true, true, 'M'},
		{21, true, false, true, 'M'},
		{20, false, true, false, 'F'}}


	e := Stuff(&friends)
	Unstuff(&comrades, e)

	if len(friends) != len(comrades) {
		t.Fatal("Nope, len.")
	}
	if cap(friends) != cap(comrades) {
		t.Fatal("Nope, cap.")
	}

	for i := range friends {
		if friends[i] != comrades[i] {
			t.Fatal("Nope.")
		}
	}

}
