package main

import "log"

func main() {
	a := removeDuplicates([]int{1, 1, 2})
	println(a)
}

func removeDuplicates(nums []int) int {
	defer log.Println(nums)
	index := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[index] {
			index++
			nums[index] = nums[i]
		}
	}
	return index + 1
}
