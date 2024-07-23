package main

import "log"

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}

/*
func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	var temp = make([]int, m+n)
	c, j := 0, 0
	for i := 0; i < m+n; i++ {
		if j == n {
			temp[i] = nums1[c]
			c++
			continue
		}
		if nums1[c] < nums2[j] && c < m {
			temp[i] = nums1[c]
			c++
			log.Println("c ", c)
		} else {
			temp[i] = nums2[j]
			j++
			log.Println("j ", j)
		}
		log.Println(temp)
	}
	copy(nums1, temp)
	log.Println(nums1)
}
*/

func merge(nums1 []int, m int, nums2 []int, n int) {
	defer log.Println(nums1)
	c, j := m-1, n-1
	for i := m + n - 1; i >= 0; i-- {
		if j < 0 {
			return
		}
		if c >= 0 && nums1[c] > nums2[j] {
			nums1[i] = nums1[c]
			c--
		} else {
			nums1[i] = nums2[j]
			j--
		}
	}
}
