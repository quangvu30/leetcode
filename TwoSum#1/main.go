package main

func main() {
	var nums = []int{2, 7, 11, 15}
	var target = 9
	var res = twoSum(nums, target)
	println(res[0], res[1])
}

func twoSum(nums []int, target int) []int {
	var maps = make(map[int]int)
	for i, v := range nums {
		if j, ok := maps[target-v]; ok {
			return []int{j, i}
		}
		maps[v] = i
	}
	return nil
}
