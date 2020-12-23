package algorithms

/*rootNode := algorithms.BSTfromSlice([]int{41,37,44,24,39,42,48,1,35,38,40,43,46,49,0,2,30,36,45,47,4,29,32,3,9,26,31,34,7,11,25,27,33,6,8,10,16,28,5,15,19,12,18,20,13,17,22,14,21,23})
	iterator := Constructor(rootNode)
	fmt.Println("Iterator list : ", iterator.list)
	fmt.Println("Iterator lenght of list : ", len(iterator.list))

	for iterator.HasNext() {
		fmt.Println("Iterator hasNext : ", iterator.HasNext())
		fmt.Println("Iterator Next : ", iterator.Next())
	}*/

type BSTIterator struct {
	list        []int
	current     int
}

func Constructor(root *TreeNode) BSTIterator {
	iterator := BSTIterator{list: []int{}, current: 0}
	iterator.InorderTraversal(root)
	return iterator
}


func (this *BSTIterator) Next() int {
	result := this.list[this.current];
	this.current++
	return result
}

func (this *BSTIterator) HasNext() bool {
	if this.current >= len(this.list) {
		return false
	}
	return true
}

func (this *BSTIterator) InorderTraversal(root *TreeNode) {

	if (root.Left != nil) && (root.Left.Val >= 0) {
		this.InorderTraversal(root.Left)
	}

	if (root.Val>=0) {
		this.list = append(this.list, root.Val)
		root.Val = -1

		if (root.Right != nil) && (root.Right.Val >= 0) {
			this.InorderTraversal(root.Right)
		}
	}
}

