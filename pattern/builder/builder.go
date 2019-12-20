package main

import (
	"fmt"
)

type Structure struct {
	Id    int
	Field string
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
}
