package main

import "testing"

func TestQuickSort(t *testing.T) {
	testTable := []struct {
		arr      []int
		expected []int
	}{
		{
			arr:      []int{-5, -9, 34, 2, 7, 23},
			expected: []int{-9, -5, 2, 7, 23, 34},
		},
		{
			arr:      []int{},
			expected: []int{},
		},
		{
			arr:      []int{4},
			expected: []int{4},
		},
		{
			arr:      []int{62, 5},
			expected: []int{5, 62},
		},
		{
			arr:      []int{53, 545, 2, 3, 8, 1, 7, 9, 34, 15, 2, 23, 43, 19},
			expected: []int{1, 2, 2, 3, 7, 8, 9, 15, 19, 23, 34, 43, 53, 545},
		},
	}

	for _, tests := range testTable {
		result := QuickSort(tests.arr)

		t.Logf("Calling: %v. Result: %v", tests.arr, result)

		if !Equal(result, tests.expected) {
			t.Errorf("Incorrect result. Expected: %v, got: %v", tests.expected, result)
		}
	}

}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
