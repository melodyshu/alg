package main

import "fmt"

/*
题目描述
一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。
请写程序找出这两个只出现一次的数字。

要求:时间复杂度是O(n)，空间复杂度是O(1)。

算法分析
因为题目限制了O(n)的时间复杂度，所以排序或者map去重不能用来解决本题。考虑到任何一个数字异或自身得到的是0，可以通过异或的方式去掉数组中重复的数字。
具体思路：异或数组中所有数字，最终得到的结果xorRes等于两个singleNumber异或的值，所以可以通过按位与求出xorRes中右边第一个等于1的二进制位。
根据该二进制位将原数组分组，在该二进制位上等于1的为一组，等于0的为另一组。易知重复的数字一定会在同一组，不同的数字一定会在不同的组，
因此可以对两个组分别异或求和，最终得到的就是两个singleNumber。

时间复杂度：O(n)，需要遍历数组
空间复杂度：O(1)
*/

func singleNumbers(nums []int) []int {
	xorRes, count := 0, 1
	x, y := 0, 0
	//数组里的所有值进行异或，结果就是只出现1次的2个数字的异或，其他的出现2次的异或值为0
	for _, v := range nums {
		xorRes ^= v
	}
	fmt.Println(xorRes)
	for xorRes & count == 0 {
		count <<= 1
	}
	for _, v := range nums {
		if v & count != 0 {
			x ^= v
		} else {
			y ^= v
		}
	}
	return []int{x, y}
}

func main() {
	arr := []int{8,7,6,7,8,1,5,6,1,2}
	fmt.Println(singleNumbers(arr))
}
