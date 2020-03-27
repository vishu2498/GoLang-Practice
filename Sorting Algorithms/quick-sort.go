package main

import "fmt"

func main() {
	arr := []int{40, 85, 69, 21, 37, 50, 19, 20}
	QuickSort(arr,0,len(arr)-1)
	fmt.Println(arr)
}

func QuickSort(arr []int, low, high int) {
	if low < high {
		QuickSort(arr,low,Partition(arr,low,high))
		QuickSort(arr,Partition(arr,low,high)+1,high)
	}
}

func Partition(arr []int, low, high int) int {
	pivot := arr[low]
	i := low
	j := high

	for i < j {
		for arr[i] <= pivot {
			i++
		}
		for arr[j] > pivot {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[low], arr[j] = arr[j], arr[low]
	return j
}
