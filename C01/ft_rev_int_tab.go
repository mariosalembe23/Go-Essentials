package main

import "fmt"

func ft_rev_int_tab(tab *[]int, size int) {
	end := size - 1

	for start := 0; start < end; start, end = start+1, end-1 {
		temp := (*tab)[start]
		(*tab)[start] = (*tab)[end]
		(*tab)[end] = temp
	}
}