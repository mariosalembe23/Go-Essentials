package main

import "fmt"

func ft_swap(a *int, b *int) {
	temp := *a
	*a = *b;
	*b = temp
}