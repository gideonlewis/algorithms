package main

import "fmt"

func main() {
	arr := []int{11, 8, 22, 2, 19, 43, 28, 38}
	/* 1. Selection sort */
	arr1 := arr
	selectionSort(arr1)

	/* 2. Bubble sort */
	arr2 := arr
	bubbleSort(arr2)

	/* 3. Insertion sort */
	arr3 := arr
	insertionSort(arr3)

	/* 4. Merge sort */
	arr4 := arr
	arr4 = mergeSort(arr4)
	printArray(arr4)

	/* 5. Quick sort */
	arr5 := arr
	quickSort(arr5, 0, len(arr5)-1)
	printArray(arr5)

	/* 6. Heap sort */
	arr6 := arr
	heapSort(arr6)
	printArray(arr6)
}

func heapSort(arr []int) {
	n := len(arr)

	// Build max heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// Extract elements from the heap one by one
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0] // Swap the root (max element) with the last element
		heapify(arr, i, 0)              // Max heapify the reduced heap
	}
}

func heapify(arr []int, n, i int) {
	largest := i     // Initialize the largest element as the root
	left := 2*i + 1  // Left child
	right := 2*i + 2 // Right child

	// If the left child is larger than the root
	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	// If the right child is larger than the largest so far
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// If the largest element is not the root
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i] // Swap the root with the largest element
		heapify(arr, n, largest)                    // Recursively heapify the affected sub-tree
	}
}

func quickSort(arr []int, low, high int) {
	if low < high {
		// Partition the array and get the pivot index
		pivotIndex := partition(arr, low, high)

		// Recursively sort the left and right subarrays
		quickSort(arr, low, pivotIndex-1)
		quickSort(arr, pivotIndex+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high] // Choose the last element as the pivot
	i := low - 1

	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i] // Swap arr[i] and arr[j]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1] // Swap arr[i+1] and arr[high]
	return i + 1                              // Return the pivot index
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0)
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}

func insertionSort(arr []int) {
	n := len(arr)
	swapped := false
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
			swapped = true
		}
		arr[j+1] = key
	}

	if !swapped {
		fmt.Println("Array is sorted")
	} else {
		printArray(arr)
	}
}
func bubbleSort(arr []int) {
	n := len(arr)
	swapped := false
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// swap
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
	}

	if !swapped {
		fmt.Println("Array is sorted")
	} else {
		printArray(arr)
	}
}

func selectionSort(arr []int) {
	n := len(arr)
	swapped := false

	for i := 0; i < n; i++ {
		minIndex := i
		// Find the index of minimum element in the array unsorted.
		for j := i + 1; j < n; j++ {
			// If arr[minIndex] > arr[j] -> minIndex = j
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		// Swap the minimum element with the arr[minIndex] and arr[i]
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}

	if !swapped {
		fmt.Println("Array is sorted")
	} else {
		printArray(arr)
	}
}

func printArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
