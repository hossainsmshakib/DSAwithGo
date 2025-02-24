package main

import "fmt"

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func main() {
	arr := []int{12, 11, 13, 5, 6}
	insertionSort(arr)
	fmt.Println("Sorted Array:", arr)
}

/*
💡 Complexity Analysis
Case		Time Complexity		Explanation
Best Case	O(n)				If the array is already sorted, only n-1 comparisons happen, no shifting.
Worst Case	O(n²)				If the array is reverse sorted, every element must be compared and shifted.
Average Case	O(n²)			Random order: On average, elements are shifted n/2 times.
Space Complexity	O(1)		In-place sorting, no extra memory is used.
*/
