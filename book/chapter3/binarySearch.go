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
