package main

import (
	"generic"
)

///////////////////////////////////////////////////////////////////////// COLMGR

const colmgrRoot = ^uintptr(0)

type Collector interface {
	Atterer
}
type Atterer interface {
	At(uintptr) Atter
}
type Atter interface {
	UpdMapper
}
type UpdMapper interface {
	Map() *generic.Value
	Upd(*generic.Value)
}
type Rooter interface {
	Root() Collector
}

var collections map[generic.Container]Collector

func init() {
	collections = make(map[generic.Container]Collector)
}
func Init(np *generic.Container, rooter Rooter) {
	p := *np

	if collections[p] != nil {
		panic("Reinitialized collection")
	}

	collections[p] = rooter.Root()
}

func Destroy(np *generic.Container) {
	p := *np

	delete(collections, p)

	print("Destroyed\n")
}
func At(np *generic.Container, key uintptr) Atter {
	p := *np
	return collections[p].At(key)
}

//////////////////////////////////////////////////////////////////// COLLECTION

func (n *Node) Map() *generic.Value {
	return &n.Val
}

func (n *Node) Upd(v *generic.Value)  {
	n.Val = *v
}

type Extra struct {}

func (e Extra) Root() (Collector) {
	var r Root
	r.Extra = e
	return r
}

type Root struct {
	generic.Is
	Extra
	trunk Node
}

func (r Root) At(v uintptr) Atter {
	return &r.trunk
}

type Node struct {
	Val generic.Value
}


////////////////////////////////////////////////////////////////////////////////

func main() {
	var a *[]int
	var b *[]byte
	var c *interface{}

	Init(&a, Root{})  // We initialize the collection handle with our tree
	defer Destroy(&a) // We destroy the collection
	Init(&b, Root{})  // We initialize the collection handle with our tree
	defer Destroy(&b) // We destroy the collection
	Init(&c, Root{})  // We initialize the collection handle with our tree
	defer Destroy(&c) // We destroy the collection

	va := []int{1,3,3,7}
	vb := []byte("1337")
	var vc interface{} = 1337

	pa := At(&a, colmgrRoot)

}
