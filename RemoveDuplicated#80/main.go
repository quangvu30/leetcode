package main

import "log"

func main() {
	a := removeDuplicates([]int{0, 0, 1, 1, 1, 1, 2, 3, 3})
	println(a)
}

func removeDuplicates(nums []int) int {
	defer log.Println(nums)
	index := 0
	c := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[index] {
			index++
			nums[index] = nums[i]
			c = 1
		} else {
			c++
			if c == 2 {
				index++
				nums[index] = nums[i]
			}
		}
	}
	return index + 1
}
