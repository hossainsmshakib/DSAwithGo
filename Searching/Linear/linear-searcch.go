/*
✅ Time Complexity: (O(N)) → On average, we check about half the elements before finding the target.
✅ Space Complexity: O(1) → We use only a constant amount of extra space (no additional data structures).
*/

package main

import "fmt"

func linearSearch(arr []int, target int) int{
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	target := 3
	result := linearSearch(arr, target)
	fmt.Println(result)
}