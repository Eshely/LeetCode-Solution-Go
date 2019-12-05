// 参考：https://www.jianshu.com/p/b6ad653fb2e1

package main

import "fmt"

func search(nums []int, target int) int {
	if nums == nil { // 数组为空时，返回-1
		return -1
	}

	left := 0
	right := len(nums) - 1
	for left < right {
		mid := (right + left) >> 1 // 求中位数
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if nums[left] == target { // 判断该数是否存在
		return left
	} else {
		return -1
	}
}

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	ret := search(nums, 9)
	fmt.Println(ret)
}
