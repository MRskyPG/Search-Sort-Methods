package main

func QuickSort(arr []int) []int {
	mas := make([]int, len(arr))
	copy(mas, arr)
	doSort(mas, 0, len(arr)-1)
	return mas
}

func doSort(arr []int, first int, last int) {
	if first >= last {
		return
	}
	i := first
	j := last
	x := arr[(first+last)/2]
	for i < j {
		for arr[i] < x {
			i++
		}
		for arr[j] > x {
			j--
		}
		if i <= j {
			temp := arr[i]
			arr[i] = arr[j]
			arr[j] = temp
			i++
			j--
		}
	}
	doSort(arr, first, j)
	doSort(arr, i, last)
}
