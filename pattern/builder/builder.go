package main

import (
	"fmt"
)

type Continuation interface {
	Continue(f func(structure *Structure) Continuation) Continuation
	End() *Structure
}

type Structure struct {
	Id    int
	Field string
}

func (structure *Structure) Continue (f func(structure *Structure) Continuation) Continuation {
	return f(structure)
}

func (structure *Structure) End () *Structure {
	return structure
}

func NewStructure() *Structure {
	return &Structure{}
}

func (structure *Structure) SetId(id int) *Structure {
	structure.Id = id
	return structure
}

func (structure *Structure) SetField(field string) *Structure {
	structure.Field = field
	return structure
}

func (structure *Structure) Build() *Structure {
	return structure
}

func main() {
	// simple builder
	structure := NewStructure().SetId(1).SetField("field").Build()
	fmt.Printf("%v\n", structure)

	// なんかできた
	// upcast, downcastはinterface実装側で持たせたい気持ちがある
	// もうちょっとなんか考えられそう
	strcont := NewStructure().Continue(func (structure *Structure) Continuation {
		structure.SetId(1)
		return (interface{}(structure)).(Continuation)
	}).Continue(func (structure *Structure) Continuation {
		structure.SetField("field")
		return (interface{}(structure)).(Continuation)
	}).End()

	fmt.Printf("%v\n", strcont)
}
