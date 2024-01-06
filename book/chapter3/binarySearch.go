package chapter3

func BinarySearch(item int, arr []int, low, high int) int {
	if low > high {
		return -1
	}
	mid := low + (high-low)/2
	if arr[mid] < item {
		return BinarySearch(item, arr, mid+1, high)
	} else if arr[mid] > item {
		return BinarySearch(item, arr, low, mid-1)
	} else {
		return mid
	}

}
