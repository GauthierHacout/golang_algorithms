package algorithms

import (
	"sort"
)

// Each token has a value in power, you can sell one token for one points and gain it's power or
// buy a token for a point in exchange of it's value in power, find out how many points you can have maximum with
// initial power P
func BagOfTokensScore(tokens []int, P int) int {

	points := 0;
	power := P;

	// Sort the tokens by ascending values
	QuickSort(tokens);

	// Buy the lowest cost in power
	for {
		if !(len(tokens)!=0 && power>tokens[0]) {
			break;
		}
		power -= buyFirstToken(&tokens);
		points++;
	}

	// Repeat as much as possible
	for {

		if !(len(tokens)>1) || points==0{
			break;
		}
		// Sell the highest cost for power
		power += sellLastToken(&tokens);
		points--;

		// Buy as much as possible
		for ok := true; ok; ok = (len(tokens)!=0 && tokens[0] <= power) {
			power -= buyFirstToken(&tokens);
			points++;
		}
	}

	return points;
}

func buyFirstToken(tokens *[]int) int{
	power := (*tokens)[0];
	*tokens = (*tokens)[1:];

	return power
}

func sellLastToken(tokens *[]int) int{
	lastIndex := len(*tokens)-1;
	power := (*tokens)[lastIndex];
	*tokens = (*tokens)[:lastIndex];

	return power
}


// Simpler implementation for the same problem, found later
func BagOfTokensScoreSimpler(tokens []int, P int) int {
	sort.Ints(tokens)
	points := 0
	lo, hi := 0, len(tokens)-1
	for lo <= hi {
		if P-tokens[lo] >= 0 {
			points++
			P -= tokens[lo]
			lo++
		} else if points > 0 && lo < hi{
			points--
			P += tokens[hi]
			hi--
		} else {
			break
		}
	}
	return points
}