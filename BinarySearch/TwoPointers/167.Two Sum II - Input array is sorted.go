package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	/* 双指针
	思路：从数组两边到中间逼近，如果和大于目标值，表示右边的值过大，应该减小，反之，增大左边的值
	*/
	size := len(numbers)
	if size < 2 {
		return []int{}
	}

	left, right := 0, size-1
	for left < right {
		if numbers[left]+numbers[right] > target { // 从右往左逼近
			right--
		} else if numbers[left]+numbers[right] < target {  // 从左往右逼近
			left++
		} else {
			return []int{left + 1, right + 1}
		}
	}
	return []int{}
}

func main() {
	numbers := []int{2, 7, 11, 15}
	fmt.Println(twoSum(numbers, 9))
}

