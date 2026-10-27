func replaceElements(arr []int) []int {
	for i := 0; i < len(arr) - 1; i++ {
		arr[i] = arr[i + 1]
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > arr[i] {
				arr[i] = arr[j]
				fmt.Println("i: ", i, ", j: ", j, ", array: ", arr)
			}
		}
	}
	arr[len(arr) - 1] = -1
	return arr
}
