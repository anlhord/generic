package typeof

import (
	"github.com/anlhord/generic"
	"testing"
)

type test struct {
	is generic.Is
	pad0 uint16
	alsois generic.Is
	pad1 uint16
	col generic.Collection
	pad2 uint16
	alsocol generic.Collection
	pad3 uint16
	myinf generic.Interface
	pad4 uint16
	alsoinf generic.Interface
	pad5 uint16
	i int
	j int
}

func TestTypeof(t *testing.T) {
	var s test
	s.pad0 = 7
	s.pad1 = 8
	s.pad2 = 9
	s.pad3 = 6
	s.pad4 = 5
	s.pad5 = 4


	a := Typeof(&s.is)
	c := Typeof(&s.alsois)

	b := Typeof(&s.col)
	f := Typeof(&s.alsocol)

	d := Typeof(&s.myinf)
	e := Typeof(&s.alsoinf)

	g := Typeof(&s.i)
	h := Typeof(&s.j)

	// Is must pass
	if !Is(a,c) {
		t.Fatal("Nope is.")
	}
	if Is(b,f) {
		t.Fatal("Nope col.")
	}
	if !Is(d,e) {
		t.Fatal("Nope inf.")
	}
	if Is(g,h) {
		t.Fatal("Nope int.")
	}

	if a != c {
		t.Fatal("Nope is.")
	}
	if b == f {
		t.Fatal("Nope col.")
	}
	if g == h {
		t.Fatal("Nope int.")
	}
	if d != e {
		t.Fatal("Nope inf.")
	}
}
