package main

import "log"

func main() {
	a := candy([]int{1, 2, 87, 87, 87, 2, 1})
	println(a)
}

func candy(ratings []int) int {
	candies := make([]int, len(ratings))
	defer log.Println(candies)
	for i := 0; i < len(candies)-1; i++ {
		candies[i] = 1
		if ratings[i] > ratings[i+1] {
			candies[i] = candies[i] + 1
		}
	}
	return Sum(candies)
}

func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}
