package main

type GenericNode[ValueType any] struct {
	Parent *GenericNode[ValueType]
	Child  []*GenericNode[ValueType]
	Value  *ValueType
}

type Category struct {
	id       int
	parentId *int
}
