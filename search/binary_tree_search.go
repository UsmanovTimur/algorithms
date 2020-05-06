package search

type BinaryNode struct {
	value int
	index int
	left  interface{}
	right interface{}
}

func (node *BinaryNode) Add(value, index int) {
	if value <= node.value {
		if node.left == nil {
			new_node := BinaryNode{
				value: value,
				index: index,
			}
			node.left = new_node
		} else {
			aaa := node.left.(BinaryNode)
			aaa.Add(value, index)
			node.right = aaa
		}
	} else {
		if node.right == nil {
			new_node := BinaryNode{
				value: value,
				index: index,
			}
			node.right = new_node
		} else {
			aaa := node.right.(BinaryNode)
			aaa.Add(value, index)
			node.right = aaa
		}
	}
}
func (node *BinaryNode) FindElement(value int) int {
	if node.value == value {
		return node.index
	}
	if value > node.value {
		if node.right == nil {
			return -1
		}
		rNode := node.right.(BinaryNode)
		return rNode.FindElement(value)
	}
	if value < node.value {
		if node.left == nil {
			return -1
		}
		lNode := node.left.(BinaryNode)
		return lNode.FindElement(value)
	}
	return -1
}

func BinaryTreeSearch(arr []int, findElem int) int {
	if len(arr) == 0 {
		return -1
	}
	if len(arr) == 1 {
		if arr[0] == findElem {
			return 0
		} else {
			return -1
		}
	}
	//Формируем дерево
	tree := BinaryNode{
		value: arr[0],
		index: 0,
	}
	for i := 1; i < len(arr); i++ {
		tree.Add(arr[i], i)
	}
	//Поиск элемента
	index := tree.FindElement(findElem)
	return index
}
