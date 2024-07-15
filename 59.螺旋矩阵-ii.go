package main

/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 */

// @lc code=start
func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	num := 1
	for round := 0; round < (n+1)/2; round++ {
		if num == n*n {
			matrix[(n-1)/2][(n-1)/2] = n * n
		}
		for col := 0; col < n-1-round*2; col++ {
			matrix[round][round+col] = num
			num++
		}
		for row := 0; row < n-1-round*2; row++ {
			matrix[round+row][n-1-round] = num
			num++
		}
		for col := 0; col < n-1-round*2; col++ {
			matrix[n-1-round][n-1-round-col] = num
			num++
		}
		for row := 0; row < n-1-round*2; row++ {
			matrix[n-1-round-row][round] = num
			num++
		}
	}
	return matrix
}

// @lc code=end
