package main

import "fmt"

func main() {
	nums := []int{32, 15, 44, 11, 2}
	quickSort(nums, 0, len(nums)-1)
	fmt.Printf("result: %v\n", nums)
}

func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	} else {
		position := partition(nums, l, r)
		quickSort(nums, l, position-1)
		quickSort(nums, position+1, r)
	}
}

func partition(nums []int, l, r int) int {
	mid := nums[l]
	first := l + 1
	for i := l + 1; i <= r; i++ {
		if nums[i] < mid {
			nums[first], nums[i] = nums[i], nums[first]
			first++
		}
	}
	nums[l] = nums[first-1]
	nums[first-1] = mid
	return first - 1
}
