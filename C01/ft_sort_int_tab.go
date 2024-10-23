package main

import "fmt"

func ft_sort_int_tab(tab *[]int, size int) {
	arr := *tab
	for i := 0; i < size - 1; i++ {
		for i := 0; i < size - 1; i++ {
			if arr[i] > arr[i + 1] {
				temp := arr[i]
				arr[i] = arr[i + 1]
				arr[i + 1] = temp
			}
		}
	}
}