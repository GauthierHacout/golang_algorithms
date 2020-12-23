package algorithms

import (
	"math"
	"sort"
)

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func makeArrayIncreasing1(a []int, b []int) int {
	sort.Ints(b)    // step 1
	m := 0
	for _, v := range b {     // step 1
		if m == 0 || v != b[m-1] {
			b[m] = v
			m++
		}
	}
	b = b[:m]
	// fmt.Println(b)
	n := len(a)

	f := make([]int, m+1) // f[m] = 0
	for i := 0; i < m; i++ {
		f[i] = 1
	}
	g := make([]int, m+1)

	for i := 1; i < n; i++ {
		// step 3
		if a[i-1] >= a[i] {
			g[m] = 1 << 30
		} else {
			g[m] = f[m]
		}
		for j := 0; j < m; j++ {
			if b[j] < a[i] {
				g[m] = min(g[m], f[j])
			}
		}

		// step 4
		for j := 0; j < m; j++ {
			g[j] = 1 << 30
			if j > 0 {
				g[j] = min(g[j], f[j-1]+1)  // step 5
			}
			if a[i-1] < b[j] {
				g[j] = min(g[j], f[m]+1)
			}

		}
		// fmt.Println(g)
		f, g = g, f   // step 6
	}
	ans := 1 << 30
	for _, v := range f {
		ans = min(ans, v)
	}

	if ans > n {
		return -1
	} else {
		return ans
	}
}


func makeArrayIncreasing(arr1 []int, arr2 []int) int {
	arr2 = sortAndDedup(arr2)
	m, n := len(arr1), len(arr2)
	dp := [2][]int{}
	for i := range dp {
		dp[i] = make([]int, n+1)
		if i == 0 {
			for j := range dp[i] {
				dp[i][j] = math.MaxInt64
			}
		}
	}
	minOps := math.MaxInt64
	for i := m-1; i >= 0; i-- {
		for j := n; j >= 0; j-- {
			var prev int
			if i == 0 {
				prev = math.MinInt64
			} else if j == n {
				prev = arr1[i-1]
			} else {
				prev = arr2[j]
			}
			k := ceiling(arr2, prev)
			dp[0][j] = math.MaxInt64
			if arr1[i] > prev {
				dp[0][j] = dp[1][n]
			}
			if k < n && dp[1][k] != math.MaxInt64 {
				dp[0][j] = min(dp[0][j], 1+dp[1][k])
			}
			if i == 0 {
				minOps = min(minOps, dp[0][j])
			}
		}
		dp[0], dp[1] = dp[1], dp[0]
	}
	if minOps == math.MaxInt64 {
		return -1
	}
	return minOps
}

func ceiling(A []int, target int) int {
	n := len(A)
	l, r := 0, n
	for l < r {
		m := l+(r-l)/2
		if A[m] > target {
			r = m
		} else {
			l = m+1
		}
	}
	return l
}

func sortAndDedup(A []int) []int {
	sort.Ints(A)
	w, n := 0, len(A)
	for i := range A {
		A[w] = A[i]
		w++
		for i+1 < n && A[i+1] == A[i] {
			i++
		}
	}
	return A[:w]
}
