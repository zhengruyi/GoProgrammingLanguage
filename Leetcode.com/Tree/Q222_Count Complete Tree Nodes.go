package Tree


func countNodes(root *TreeNode) int {
	h := height(root)
	if h < 0 {
		return 0
	}else {
		if height(root.Right) == h-1 {
			//表示root的左子树完整,那么左子树一共有2^(h-1) -1 ,加上根节点叫 2^(h-1)
			return (1<< h) + countNodes(root.Right)
		}else{
			//表示缺失节点在左子树上,那么右子树是完整的,高度为(h-2),加上根节点总数为2^(h-2)
			return (1 << (h-1)) + countNodes(root.Left)
		}
	}

}

//求树的高度,对于root == nil 返回-1
func height(root *TreeNode) int {
	if root == nil {
		return -1
	}else{
		return 1 + height(root.Left)
	}
}
