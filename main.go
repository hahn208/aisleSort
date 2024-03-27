package main

import "fmt"

// Given two ints and the current iteration status flag, compare ints and return in sorted order.
func doSort(k, l int, isSwapped bool) (int, int, bool) {
	if k > l {
		return l, k, true
	}
	return k, l, isSwapped
}

func aisleSort(toSort []int) []int {
	isSwapped := true
	start := 0
	end := len(toSort) - 1

	for isSwapped {
		isSwapped = false

		// Process from tip to tail.
		for i := start; i < end; i++ {
			// If the ints are not in order, they will be returned in order with a bool True flag.
			toSort[i], toSort[i+1], isSwapped = doSort(toSort[i], toSort[i+1], isSwapped)
		}

		// No swaps were performed, all sorted.
		if !isSwapped {
			break
		}

		// Update the upper limit because the last seat is highest int.
		end--

		// Process tail to tip.
		for i := end - 1; i >= start; i-- {
			toSort[i], toSort[i+1], isSwapped = doSort(toSort[i], toSort[i+1], isSwapped)
		}

		// Front seat has the lowest int. Set the lower limit index.
		start++
	}

	return toSort
}

func main() {
	unsorted := []int{4, 6, 3, 5, 1, 4, 5, 3, 9}

	fmt.Println(aisleSort(unsorted))
}
