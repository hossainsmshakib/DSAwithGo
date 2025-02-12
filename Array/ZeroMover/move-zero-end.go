/*
✅ O(N) Time Complexity → Each element is processed once, and swaps happen at most N times.
✅ O(1) Space Complexity → No extra array; everything happens in-place.
*/

package main

import "fmt"

func moveZero(arr []int) {
	// Pointer to track the position of the next non-zero element
	left := 0

	for right := 0; right < len(arr); right++ {
		if arr[right] != 0 {
			arr[left], arr[right] = arr[right], arr[left]
			left++
		}
		// Print visualization of each step
		fmt.Printf("Step %d: arr = %v, left = %d\n", right+1, arr, left)
	}
}

func main() {
    arr := []int{0, 1, 0, 3, 12}
    fmt.Println("Original Array:", arr)

    moveZero(arr)

    fmt.Println("Final Result:", arr)
}