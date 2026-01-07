package main

func nFibonacci(n int) []int {
	if n <= 0 {
		return []int{}
	}

	if n == 1 {
		return []int{0}
	}

	if n == 2 {
		return []int{0, 1}
	}

	ans := make([]int, n)
	ans[0] = 0
	ans[1] = 1
	for i := 2; i < n; i++ {
		ans[i] = ans[i-1] + ans[i-2]
	}

	return ans
}
