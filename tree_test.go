package main

import (
	"testing"

	"gotest.tools/assert"
)

type testCase struct {
	Annotation string
	Input      []Category
	Expected   map[int]*GenericNode[Category]
	HasError   bool
}

var node0 = GenericNode[Category]{
	Child: []*GenericNode[Category]{},
	Value: &Category{
		id: 0,
	},
}

var node1 = GenericNode[Category]{
	Child: []*GenericNode[Category]{},
	Value: &Category{
		id:       1,
		parentId: Ptr(0),
	},
	Parent: &node0,
}

var node2 = GenericNode[Category]{
	Value: &Category{
		id:       2,
		parentId: Ptr(0),
	},
	Parent: &node0,
}

var node3 = GenericNode[Category]{
	Value: &Category{
		id:       3,
		parentId: Ptr(1),
	},
	Parent: &node1,
}

func TestMakeTree(t *testing.T) {
	node0.Child = append(node0.Child, &node1)
	node0.Child = append(node0.Child, &node2)
	node1.Child = append(node1.Child, &node3)

	output1 := map[int]*GenericNode[Category]{
		0: &node0,
		1: &node1,
		2: &node2,
		3: &node3,
	}

	testCases := []testCase{
		{
			Annotation: "Legit MakeTree",
			Input: []Category{
				{
					id: 0,
				},
				{
					id:       1,
					parentId: Ptr(0),
				},
				{
					id:       2,
					parentId: Ptr(0),
				},
				{
					id:       3,
					parentId: Ptr(1),
				},
			},
			Expected: output1,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Expected
		result := MakeTree(testCase.Input)

		assert.Equal(t, actual, result, testCase.Annotation)
	}
}

func TestFilterTree(t *testing.T) {
	node0.Child = append(node0.Child, &node1)
	node0.Child = append(node0.Child, &node2)
	node1.Child = append(node1.Child, &node3)

	output1 := map[int]*GenericNode[Category]{
		0: &node0,
		3: &node3,
	}

	testCases := []testCase{
		{
			Annotation: "Legit MakeTree",
			Input: []Category{
				{
					id: 0,
				},
				{
					id:       1,
					parentId: Ptr(0),
				},
				{
					id:       2,
					parentId: Ptr(0),
				},
				{
					id:       3,
					parentId: Ptr(1),
				},
			},
			Expected: output1,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Expected
		treeMap := MakeTree(testCase.Input)
		result := FilterTreeByParentIds([]int{0, 3}, treeMap)

		assert.Equal(t, actual, result, testCase.Annotation)
	}
}
