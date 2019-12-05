package main

import "fmt"

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}

	//x的二分之一必定小于等于x的开方。右边界取x的二分之一就能使查找区间更小，程序更有效率
	left, right := 1, x>>1 // 等价于 x / 2
	for left < right {
		mid := left + ((right - left + 1) >> 1) // 取右中位数
		if mid*mid > x {
			right = mid - 1
		} else {
			left = mid
		}
	}
	return left
}

func main() {
	x := 123
	fmt.Println(mySqrt(x))
}
