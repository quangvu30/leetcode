package main

import "sort"

func main() {
	var num = 3749
	var res = ""
	var maps = map[int]string{900: "CM", 500: "D", 400: "CD", 100: "C", 90: "XC", 50: "L", 40: "XL", 10: "X", 9: "IX", 5: "V", 4: "IV", 1: "I"}
	if num >= 1000 {
		temp := num / 1000
		num -= temp * 1000
		for i := 0; i < temp; i++ {
			res += "M"
		}
	}
	keys := make([]int, 0)
	for k, _ := range maps {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for i := len(keys) - 1; i >= 0; i-- {
		for num >= keys[i] {
			num -= keys[i]
			res += maps[keys[i]]
		}
		println(num, keys[i], res)
	}
	println(res)
}
