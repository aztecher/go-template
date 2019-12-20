package main

// Implementing Builder Pattern with call method constraint using Interface and tagless final

import (
	"fmt"
)

// abstract the builder pattern's base structure
type Continuation interface {
	Continue(f func(structureSetter StructureSetter) *Structure) Continuation
	Build() *Structure
}

// Constraint of call method
type StructureSetter interface {
	SetId(id int) *Structure
	SetField(field string) *Structure
}

// target for build
type Structure struct {
	Id    int
	Field string
}

func (structure *Structure) Continue (f func(strSetter StructureSetter) *Structure) Continuation {
	str := f(interface{}(structure).(StructureSetter))
	return (interface{}(str)).(Continuation)
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

func (structure *Structure) TestMethod() *Structure{
	return nil
}

func main() {
	// simple builder
	structure := NewStructure().SetId(1).SetField("field").Build()
	// if you call TestMethod(), segumentation fault. but compiler don't catch it. like bellow
	// structure := NewStructure().SetId(1).TestMethod().SetField("field").Build()
	fmt.Printf("%v\n", structure)

	// builder with call method constraint
	// TestMethod() can't suggest and if you call it then compiler caught it as error.
	strcont := NewStructure().Continue(func (structureSetter StructureSetter) *Structure{
		return structureSetter.SetId(1)
	}).Continue(func (structureSetter StructureSetter) *Structure {
		return structure.SetField("field")
	}).Build()

	fmt.Printf("%v\n", strcont)
}
