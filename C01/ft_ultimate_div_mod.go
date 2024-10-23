package main

import "fmt"

func ft_ultimate_div_mod(a *int, b *int) {

	if *b == 0 {
		fmt.Printf("Error: b most be different than 0")
		return 
	}

	div := *a / *b
	mod := *a % *b

	*a = div
	*b = mod
}
