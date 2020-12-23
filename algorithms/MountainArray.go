package algorithms

fmt.Println("SHOULD BE TRUE : ", validMountainArray([]int{1,3,2}) )
fmt.Println("SHOULD BE FALSE : ", validMountainArray([]int{1,2,4,5,6,8}) )

func validMountainArray(arr []int) bool {

	var variation int
	var index int
	curve := 1

	if len(arr)<3 || arr[1]<arr[0] {
		return false
	}

	for variation<2 {
		if index==len(arr) {
			break
		}

		if index>0 {
			if arr[index]==arr[index-1] {
				return false
			} else if changeOfCurve(arr[index], arr[index-1], curve) {
				curve = curve*-1
				variation++
			}
		}

		index++
	}

	return variation==1
}

func changeOfCurve(current int, previous int, curve int) bool {
	// sign >0 if increasing && <0 if decreasing
	sign := current-previous

	// curve =1 for increasing && =-1 for decreasing
	if sign*curve>0 {
		return false
	}

	return true
}

