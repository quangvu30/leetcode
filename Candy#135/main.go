package main

import (
	"log"
)

func main() {
	candy([]int{1, 2, 87, 87, 87, 2, 1})
	println(Sum([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

// func candy(ratings []int) int {
// 	candies := make([]int, len(ratings))
// 	defer log.Println(candies)
// 	candies[0] = 1
// 	for i := 1; i < len(candies); i++ {
// 		if ratings[i] > ratings[i-1] {
// 			candies[i] = candies[i-1] + 1
// 		} else if ratings[i] == ratings[i-1] {
// 			candies[i] = 1
// 		} else {
// 			candies[i] = 1
// 			if candies[i-1] <= candies[i] {
// 				candies[i-1] = candies[i] + 1
// 			}
// 			for j := i - 1; j > 0; j-- {
// 				if ratings[j] < ratings[j-1] && candies[j] >= candies[j-1] {
// 					candies[j-1] = candies[j] + 1
// 				} else {
// 					break
// 				}
// 			}
// 		}
// 	}
// 	return Sum(candies)
// }

func candy(ratings []int) int {
	candies := make([]int, len(ratings))
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	log.Println(candies)

	for i := len(ratings) - 1; i > 0; i-- {
		if ratings[i] < ratings[i-1] {
			p1 := candies[i] + 1
			if p1 > candies[i-1] {
				candies[i-1] = p1
			}
		}
	}
	log.Println(candies)

	result := len(ratings)

	for _, val := range candies {
		result = result + val
	}

	return result
}

func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}
