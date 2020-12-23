package algorithms

func CanReach(arr []int, start int) bool {
	return recursiveCanReach(arrayToList(arr), start)
}

type Node struct {
	Val int
	Explored bool
}

func arrayToList(arr []int) *[]Node {
	var result []Node
	for _,v := range arr {
		result = append(result, Node{Val: v, Explored:false})
	}
	return &result
}

func recursiveCanReach(list *[]Node, current int) bool {
	if 0<=current && current<len(*list)-1 && (*list)[current].Explored==false {
		if (*list)[current].Val == 0 {
			return true
		}
		(*list)[current].Explored = true
		return recursiveCanReach(list, current+(*list)[current].Val) || recursiveCanReach(list, current-(*list)[current].Val)
	}
	return false
}
