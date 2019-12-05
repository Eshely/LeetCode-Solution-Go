package main

import "fmt"

func searchInsert(nums []int, target int) int {
	if nums == nil {
		return -1
	}

	// 插入值有可能大于数组最右边的值（数组最大值），因此右边界应该取数组长度
	left, right := 0, len(nums)
	for left < right {
		mid := left + ((right - left) >> 1)
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func main() {
	nums := []int{1, 3, 5, 6}
	ret := searchInsert(nums, 2)
	fmt.Println(ret)
}
