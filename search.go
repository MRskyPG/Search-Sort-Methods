package main

func BinarySearch(arr []int, elem int) int {
	fst := 0
	lst := len(arr)
	for fst != lst {
		centralIdx := (fst + lst) / 2
		if elem == arr[centralIdx] {
			return centralIdx
		}
		if elem < arr[centralIdx] {
			lst = centralIdx
		} else {
			fst = centralIdx + 1
		}
	}
	return -1
}

func LinearSearch(arr []int, elem int) int {
	fst := 0
	lst := len(arr)
	for fst < lst {
		if arr[fst] == elem {
			return fst
		}
		fst++
	}
	return -1
}
