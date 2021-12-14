package main

import "fmt"

/*
题目描述
在一个 m*n 的棋盘的每一格都放有一个礼物，
每个礼物都有一定的价值（价值大于 0）。你可以从棋盘的左上角开始拿格子里的礼物，
并每次向右或者向下移动一格、直到到达棋盘的右下角。给定一个棋盘及其上面的礼物的价值，
请计算你最多能拿到多少价值的礼物？
输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 12
解释: 路径 1→3→5→2→1 可以拿到最多价值的礼物
*/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxValue(arr [][]int) int {
	//最上面一行
	for j := 1; j < len(arr[0]); j++ {
		arr[0][j] = arr[0][j-1] + arr[0][j]
	}
	//最左边一行
	for i := 1; i < len(arr); i++ {
		arr[i][0] = arr[i-1][0] + arr[i][0]
	}
	for i := 1; i < len(arr); i++ {
		for j := 1; j < len(arr[0]); j++ {
			arr[i][j] = max(arr[i-1][j], arr[i][j-1]) + arr[i][j]
		}
	}
	//最右下角的下标存储的即是最大值
	//缺点:不能打印走过的最优路径
	return arr[len(arr)-1][len(arr[0])-1]
}

func main() {
	arr := [][]int{{1, 3, 1,2}, {1, 5, 1,3}, {4, 2, 1,4}}
	fmt.Printf("最大值:%d",maxValue(arr))
}
