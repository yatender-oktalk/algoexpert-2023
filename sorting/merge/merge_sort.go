package merge

func MergeSort(arr []int) []int {
	if len(arr) > 1 {
		firstHalf := MergeSort(arr[:len(arr)/2])
		secondHalf := MergeSort(arr[len(arr)/2:])
		return Merge(firstHalf, secondHalf)
	} else {
		return arr
	}
}

func Merge(firstHalf, secondHalf []int) []int {
	sortedArr := []int{}
	i := 0
	j := 0

	for i < len(firstHalf) && j < len(secondHalf) {
		if firstHalf[i] <= secondHalf[j] {
			sortedArr = append(sortedArr, firstHalf[i])
			i++
		} else {
			sortedArr = append(sortedArr, secondHalf[j])
			j++
		}
	}

	for i < len(firstHalf) {
		sortedArr = append(sortedArr, firstHalf[i])
		i++
	}

	for j < len(secondHalf) {
		sortedArr = append(sortedArr, secondHalf[j])
		j++
	}
	return sortedArr
}
