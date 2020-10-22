package algorithms

import (
	"math"
	"strconv"
)

func ReverseInt(x int) int {
	if math.Abs(float64(x))<10 {
		return x
	}

	var sign = 0
	if x >= 0 {
		sign = 1
	} else {
		sign = -1
	}

	// You can use the mathematical method or use a string ( a string is much faster on average )

	//reversedXAsSlice := intReversedToSlice(x*sign)
	//reversedX := sliceToInt(reversedXAsSlice)

	xAsString := strconv.Itoa(x*sign)
	reversedString := reverseString(xAsString)
	reversedX, _ := strconv.Atoi(reversedString)

	if reversedX < int(math.Pow(2, 31))  {
		return sign*reversedX
	}

	return 0
}

func reverseString(x string) string {
	rns := []rune(x)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}

func intReversedToSlice(x int) []int{
	var result []int
	for i := 10; i <= x*10; i=i*10 {
		result = append(result, (x%i)/(i/10))
	}
	return result
}

func sliceToInt(s []int) int {
	var result int
	for i:=0; i<len(s); i++{
		result += s[i]*int(math.Pow10(len(s)-1-i))
	}
	return result
}

