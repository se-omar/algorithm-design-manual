package chapter3

func BinarySearch(arr []int, item int, low, high int) int {
	if low >= high {
		return -1
	}

	mid := low + (high-low)/2

	if arr[mid] < item {
		return BinarySearch(arr, item, mid+1, high)
	}

	if arr[mid] > item {
		return BinarySearch(arr, item, low, mid-1)
	} else {
		return mid
	}
}

func BinarySearchIter(arr []int, item int) int {
	low, high := 0, len(arr)
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] < item {
			low = mid + 1
		} else if arr[mid] > item {
			high = mid - 1
		} else {
			return mid
		}
	}

	return -1
}
