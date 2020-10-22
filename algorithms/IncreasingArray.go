package algorithms

import (
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
