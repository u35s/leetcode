package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if v, ok := m[nums[i]]; ok {
			return []int{v, i}
		}
		x := target - nums[i]
		m[x] = i
	}
	return []int{0, 0}
}

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 5))
}
