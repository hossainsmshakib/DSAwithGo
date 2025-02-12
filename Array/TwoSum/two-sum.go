// Time - O(N*N)
// Space - O(N)

package main

import "fmt"

func twoSum(arr []int, target int) []int{
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func main() {
	arr := []int{1, 2, 2, 4, 6}
	target := 7
	result := twoSum(arr, target)
	fmt.Println("Result is: ", result)
}