package selection

func SelectionSort(arr []int) []int {
	for current_index, _ := range arr {
		arr = insert(arr, current_index)
	}
	return arr
}

func insert(arr []int, current_index int) []int {
	// first we will see all values which are more than current_index including the current_index
	// start from current_index and end at the length(arr) once it fetches the smallest one, we will swap and break the loop
	smallest_index := current_index
	for i := current_index; i < len(arr); i++ {
		if arr[i] < arr[smallest_index] {
			// save the index of the smallest value yet, iterate over all the items and find which index has the minimun value
			smallest_index = i
		}
	}
	arr[current_index], arr[smallest_index] = arr[smallest_index], arr[current_index]
	return arr
}
