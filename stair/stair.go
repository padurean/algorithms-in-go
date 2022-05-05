package stair

func Solution(n int64) int64 {
	if n == 0 {
		return 1
	}
	if n < 3 {
		return n
	}
	if n == 3 {
		return 4
	}

	// first three stairs, the loop index and result
	var a, b, c, i, r int64

	// hardcoded results for the first 3 stairs
	a = 1
	b = 2
	c = 4

	// starting from 4 as thr first 3 stairs are hardcoded
	for i = 4; i <= n; i++ {
		r = c + b + a
		a = b
		b = c
		c = r
	}

	return r
}
