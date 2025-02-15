package main

import "fmt"

func bubbleSort(arr []int) {
	for i := 0;   i < len(arr)-1; i++ {
		swapped := false
		for j := 0;   j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func main() {
	arr := []int{2, 3, 5, 1, 4, 12, 6}
	bubbleSort(arr)
	fmt.Println("Sorted array is: ", arr)
}