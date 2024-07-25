package main

import "log"

func main() {
	a := trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
	b := trap([]int{3, 2, 3, 1, 2, 3, 2, 0, 1, 2, 1, 2})
	log.Println(a, b)
}

func trap(height []int) int {
	sum := 0
	start := -1
	for i := 0; i < len(height); i++ {
		if height[i] != 0 && start == -1 {
			start = i
			continue
		} else {
			sum += i - start - 1
			start = -1
		}
	}
	return sum
}

func max(arr []int) int {
	max := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}
