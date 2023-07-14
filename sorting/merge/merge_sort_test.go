package merge

import "testing"

func TestMergeSorting(t *testing.T) {
	want := []int{1, 1, 2, 3, 3, 4, 5}
	got := MergeSort([]int{5, 4, 3, 2, 1, 3, 1})
	if len(got) != len(want) {
		t.Errorf("got %v want %v", got, want)
	}

	for i := range got {
		if got[i] != want[i] {
			t.Errorf("got %v want %v", got, want)
		}
	}
}
