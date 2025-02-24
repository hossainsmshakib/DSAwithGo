package main

import "fmt"

func removeDuplicates(arr []int) []int {
	seen := make(map[int]bool)
	result := []int{}

	for _, value := range arr {
		if !seen[value] {
			seen[value] = true
			result = append(result, value)
		}
	}
	return result
}

func main() {
	arr := []int{1, 2, 2, 3, 4, 4, 5, 6, 6, 7, 8, 8, 9}
	fmt.Println("Original Array:", arr)

	uniqueArr := removeDuplicates(arr)
	fmt.Println("Array after Removing Duplicates:", uniqueArr)
}

/*
🔹 **Time Complexity:** O(n)
   - We iterate through the array **once** (O(n)).
   - Hash map operations (lookup & insert) are **O(1)** on average.
   - Overall, the complexity is **O(n)**.

🔹 **Space Complexity:** O(n)
   - We store unique elements in a new slice (`result`), which in the worst case is O(n).
   - The hash map (`seen`) also requires O(n) space for tracking unique elements.

🔹 **Is it Iterative?** ✅ YES
   - The function uses **a single loop** (iterative approach) to process the array.
   - There’s no recursion involved.

🔹 **Why is this optimal?**
   - This approach ensures that **duplicates are removed in a single pass**.
   - It **preserves order**, unlike solutions using sorting.
   - Hash map lookups are **O(1)** on average, making it faster than nested loops.

🚀 **Best for interviews!** Uses hash maps effectively and works efficiently.
*/
