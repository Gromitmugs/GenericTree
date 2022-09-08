package main

func MakeTree(input []Category) map[int]*GenericNode[Category] { // return as map of [parentID:GenericNode]
	treeMap := make(map[int]*GenericNode[Category])

	// generate nodes
	for _, e := range input {
		newNode := makeNode(e.id, e.parentId)
		treeMap[e.id] = newNode
	}

	// build a tree, connecting parent and children
	for _, node := range treeMap {
		if node.Value.parentId != nil {
			node.addParentNode(treeMap[*node.Value.parentId])
			treeMap[*node.Value.parentId].addChildNode(node)
		}
	}

	return treeMap
}

func FilterTreeByParentIds(parentIds []int, treeMap map[int]*GenericNode[Category]) map[int]*GenericNode[Category] {
	filteredTree := make(map[int]*GenericNode[Category])
	for _, e := range parentIds {
		filteredTree[e] = treeMap[e]
	}
	return filteredTree
}

func makeNode(id int, parentId *int) *GenericNode[Category] {
	node_category := Category{id: id, parentId: parentId}
	node := &GenericNode[Category]{
		Value: &node_category,
	}
	return node
}

func (node *GenericNode[Category]) addChildNode(ChildNode *GenericNode[Category]) {
	node.Child = append(node.Child, ChildNode)
}

func (node *GenericNode[Category]) addParentNode(ParentNode *GenericNode[Category]) {
	node.Parent = ParentNode
}

func Ptr[t any](input t) *t {
	return &input
}
