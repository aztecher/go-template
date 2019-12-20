package main

import (
	"fmt"
)

type Continuation interface {
	Continue(f func(structure *Structure) *Structure) Continuation
	End() *Structure
}

type Structure struct {
	Id    int
	Field string
}

func (structure *Structure) Continue (f func(structure *Structure) *Structure) Continuation {
	str := f(structure)
	return (interface{}(str)).(Continuation)
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

	// 引数に与える関数の引数と返り値が同一型を取れるように変更でき、
	// かつコード側の責務が減った(builderのsetter呼びに近い形になった)
	strcont := NewStructure().Continue(func (structure *Structure) *Structure{
		return structure.SetId(1)
	}).Continue(func (structure *Structure) *Structure {
		return structure.SetField("field")
	}).End()

	fmt.Printf("%v\n", strcont)
}
