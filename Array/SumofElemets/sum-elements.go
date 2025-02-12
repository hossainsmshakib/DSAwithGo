/*
✅ Time Complexity: O(N) → We traverse the array once.
✅ Space Complexity: O(1) → We use only a single variable (sum), no extra memory.
*/

package main

import "fmt"

func sumOfElements(arr []int) int{
	sum := 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	return sum
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("Original Array: ", arr)

	result := sumOfElements(arr)
	fmt.Println("Sum of the element is: ", result)
}