package main

func lengthOfLIS(nums []int) int {
	var tails []int
	for i := 0; i < len(nums); i++ {
		tlen := len(tails)
		if tlen == 0 {
			tails = []int{nums[i]}
			continue
		}
		if nums[i] > tails[tlen-1] {
			tails = append(tails, nums[i])
			printArray(tails)
			continue
		}

		start, end, mid := 0, tlen-1, 0
		for start != end {
			mid = start + (end-start)/2
			if nums[i] > tails[mid] {
				start = mid + 1
			} else {
				end = mid
			}
		}
		tails[end] = nums[i]
		printArray(tails)
	}
	return len(tails)
}

func main() {
	var nums = []int{10, 9, 2, 5, 3, 7, 101, 18, 103}
	println(lengthOfLIS(nums))
}

func printArray(arr []int) {
	for _, v := range arr {
		print(v, " ")
	}
	println()
}
