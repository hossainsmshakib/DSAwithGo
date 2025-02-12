package main

import "fmt"

// left -> left position 0 & right -> right position 7
func reverse(arr []int, left int, right int) {
	if left >= right {
		return
	}
	// swap
	arr[left], arr[right] = arr[right], arr[left]
	// recursive call
	reverse(arr, left+1, right-1)
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("Original Array: ", arr)
	
	// passing the array and left, right position
	reverse(arr, 0, len(arr) - 1)
	fmt.Println("Reversed Array: ", arr)
}