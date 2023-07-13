package sorting

func InsertionSort(a []int) []int {
	for i := 1; i < len(a); i++ {
		a = insert(a, i)
	}
	return a
}

func insert(a []int, j int) []int {
	for i := j; i > 0; i-- {
		// if pointer element i.e. a[j] is smaller than the a[i]
		// then we will perform swap
		if a[j] < a[i-1] {
			a[j], a[i-1] = a[i-1], a[j]
			j = i - 1
		}
	}
	return a
}
