package main

import "testing"

func TestBinarySearch(t *testing.T) {
	//Arrange
	testTable := []struct {
		arr      []int
		elem     int
		expected int
	}{
		{
			arr:      []int{-4, 1, 42, 43, 44, 55, 233},
			elem:     44,
			expected: 4,
		},
		{
			arr:      []int{-5, 7, 8, 14, 16},
			elem:     7,
			expected: 1,
		},
		{
			arr:      []int{},
			elem:     3,
			expected: -1,
		},
		{
			arr:      []int{1, 6, 8, 32, 54, 58},
			elem:     -5,
			expected: -1,
		},
	}

	//Act
	for _, tests := range testTable {
		result := BinarySearch(tests.arr, tests.elem)

		t.Logf("Calling: %v and search %d. Result: %d", tests.arr, tests.elem, result)

		//Assert
		if result != tests.expected {
			t.Errorf("Incorrect result. Expected: %d, got: %d", tests.expected, result)
		}
	}

}

func TestLinearSearch(t *testing.T) {
	testTable := []struct {
		arr      []int
		elem     int
		expected int
	}{
		{
			arr:      []int{100, 2, 32, 14, 44, 98, 9},
			elem:     44,
			expected: 4,
		},
		{
			arr:      []int{2, 7, -3, 4, 112},
			elem:     7,
			expected: 1,
		},
		{
			arr:      []int{},
			elem:     3,
			expected: -1,
		},
		{
			arr:      []int{1, 6, 8, 32, 54, 58},
			elem:     -5,
			expected: -1,
		},
	}

	for _, tests := range testTable {
		result := LinearSearch(tests.arr, tests.elem)

		t.Logf("Calling: %v and search %d. Result: %d", tests.arr, tests.elem, result)

		if result != tests.expected {
			t.Errorf("Incorrect result. Expected: %d, got: %d", tests.expected, result)
		}
	}

}
