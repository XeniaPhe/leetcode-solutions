package binarysearch_cheatsheet

// standard binary search
// return the index of the target, or -1 if it doesn't exist
// if there is more than one occurrence of the target, the returned index could be any of them depending on the distribution
func ExactMatch(slice []int, target int) int {
	for left, right := 0, len(slice) - 1; left <= right; {
		mid := left + (right - left) / 2

		if slice[mid] < target {
			left = mid + 1
		} else if slice[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}

	return -1
}

// find the first match (leftmost) of the target if one or more exist
// find the index where the target would be inserted if target doesn't exist
// find the first element greater than or equal to the target
// find lower bound (inclusive) in a range
// also referred to as bisect left
func LowerBound(slice []int, target int) int {
	left, right := 0, len(slice)

	for left < right {
		mid := left + (right - left) / 2

		if slice[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

// similar to lower bound (bisect left) but returns the index after the rightmost occurrence of the target
// find upper bound (exclusive) in a range
// also referred to as bisect right
func UpperBound(slice []int, target int) int {
	left, right := 0, len(slice)

	for left < right {
		mid := left + (right - left) / 2

		if slice[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}


/* extra: you can count occurrences of elements via (upperBound - lowerBound) */


// find the exact index of the first occurrence of the target, or -1 if it doesn't exist
// alternatively can be implemented via lowerbound and checking if the index returned is the target
// also can be implemented in O(n) time via standard binary search (exact match) and a reverse linear search
func FirstOccurrence(slice []int, target int) int {
	result := -1

	for left, right := 0, len(slice) - 1; left <= right; {
		mid := left + (right - left) / 2

		if slice[mid] < target {
			left = mid + 1
		} else if slice[mid] > target {
			right = mid - 1
		} else {
			right = mid - 1
			result = mid
		}
	}

	return result
}

// find the exact index of the last occurrence of the target, or -1 if it doesn't exist
// alternatively can be implemented via upperbound and checking if the index (index returned - 1) is the target
// also can be implemented in O(n) time via standard binary search (exact match) and a linear search
func LastOccurrence(slice []int, target int) int {
	result := -1

	for left, right := 0, len(slice) - 1; left <= right; {
		mid := left + (right - left) / 2

		if slice[mid] < target {
			left = mid + 1
		} else if slice[mid] > target {
			right = mid - 1
		} else {
			result = mid
			left = mid + 1
		}
	}

	return result
}

// find the index of the greatest element that is less than or equal to the target, or -1 if no such element exists
func Floor(slice []int, target int) int {
	result := -1

	for left, right := 0, len(slice) - 1; left <= right; {
		mid := left + (right - left) / 2

		if slice[mid] <= target {
			left = mid + 1
			result = mid
		} else {
			right = mid - 1
		}
	}

	return result
}

// find the index of the smallest element that is greater than or equal to the target, or -1 if no such element exists
func Ceil(slice []int, target int) int {
	result := -1

	for left, right := 0, len(slice) - 1; left <= right; {
		mid := left + (right - left) / 2

		if slice[mid] < target {
			left = mid + 1
		} else {
			result = mid
			right = mid - 1
		}
	}

	return result
}