package binarysearch

import "testing"

func TestBinarySearch(t *testing.T) {
	arr := []int{2, 3, 5, 15, 60, 100}

	t.Run("success case", func(t *testing.T) {
		want := 2
		got := BinarySearch(arr, 0, len(arr)-1, 5)
		if got == -1 {
			t.Error("Element not present in the list")
		}
		if got != want {
			t.Error("Index not proper")
		}
	})

	t.Run("success case 2", func(t *testing.T) {
		want := 5
		got := BinarySearch(arr, 0, len(arr)-1, 100)
		if got == -1 {
			t.Error("Element not present in the list")
		}
		if got != want {
			t.Error("Index not proper")
		}
	})

	t.Run("falining case", func(t *testing.T) {
		want := -1
		got := BinarySearch(arr, 0, len(arr)-1, -5)
		if got != want {
			t.Error("Invalid case")
		}
	})
}
