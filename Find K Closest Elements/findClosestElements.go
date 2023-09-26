package main

func main() {

}

func findClosestElements(arr []int, k int, x int) []int {
	/**
	 * Find the leftmost in the k numbers closest to x.
	 * On right side we end at the kth number from end,
	 * to ensure the left-right window will always have at least k numbers
	 * TC: O(log(N - K))
	 * SC: O(1)
	 */
	l, r := 0, len(arr)-k

	for l < r {
		mid := l + (r-l)/2
		/**
		* Instead of comparing with x during binary search,
		* we compare the closeness to x
		* than the kth number on it right
		 */
		if x-arr[mid] > arr[mid+k]-x {
			l = mid + 1
		} else {
			r = mid
		}
	}
	/**
	 * Return k numbers starting from leftmost index
	 * TC: O(K)
	 * SC: O(1)
	 */
	return arr[l : l+k]
}
