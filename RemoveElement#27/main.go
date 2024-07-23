package main

import "log"

func main() {
	a := removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)
	println(a)

}

func removeElement(nums []int, val int) int {
	c := 0
	length := len(nums)
	temp := make([]int, length)
	defer log.Println(temp)
	for i := 0; i < length; i++ {
		if nums[i] != val {
			temp[c] = nums[i]
			c++
		}
	}
	copy(nums, temp)
	return c
}
