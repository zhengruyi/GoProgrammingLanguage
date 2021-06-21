package main

type tree struct {
	value int
	left,right *tree
}

func Sort(values []int)  {
	var root *tree
	for _, v := range values {
		root = add(root ,v)
	}
	appendValues(values[:0], root)
}
/*
中序遍历对搜索二叉树会得到有序序列
 */
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}
/*
根据value和当前节点的值，选择将值插在当前节点的左边还是右边

 */
func add(t *tree, value int) *tree {
	if t == nil {
		t = new (tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	}else {
		t.right = add(t.right, value)
	}
	return t
}