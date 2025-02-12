/*
✅ Time Complexity: O(N) (Single pass through array)
✅ Space Complexity: O(1) (No extra space used)
*/

package main

import "fmt"

func minMax(arr []int)(int, int) {
	min, max := arr[0], arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
	}
	return min, max
}

func main() {
	arr := []int{1, 2, 3, 6, 0}
	fmt.Println("Original Array: ", arr)

	min, max := minMax(arr)
	fmt.Println("Min and Max values are: ", min, max)

}