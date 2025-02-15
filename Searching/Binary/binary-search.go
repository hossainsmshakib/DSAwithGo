/*
1️⃣ Time Complexity: O(log N) → Each step halves the search space.
2️⃣ Space Complexity: O(1) → No extra memory needed, unlike recursion (which uses stack space).

Only works on sorted array.
*/

package main

import "fmt"

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 7, 9}
	target := 9

	result := binarySearch(arr, target)
	fmt.Println("Found at index", result ,"using binary search")
}