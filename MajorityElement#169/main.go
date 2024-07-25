package main

import "log"

func main() {
	a := majorityElement([]int{6, 5, 5, 2, 3, 2, 2})
	println(a)
}

// func majorityElement(nums []int) int {
// 	m := make(map[int]int)
// 	defer log.Println(m)
// 	max := 0
// 	majority := 0
// 	for i := 0; i < len(nums); i++ {
// 		if val, ok := m[nums[i]]; !ok {
// 			m[nums[i]] = 1
// 		} else {
// 			m[nums[i]] = val + 1
// 		}

// 		if m[nums[i]] > max {
// 			max = m[nums[i]]
// 			majority = nums[i]
// 		}
// 	}
// 	return majority
// }

func majorityElement(nums []int) int {
	var counter int
	var candidate int

	for _, num := range nums {
		if counter == 0 {
			candidate = num
		}
		if num == candidate {
			counter++
		} else {
			counter--
		}
		log.Println("candidate", candidate, " | counter", counter)
	}

	return candidate
}
