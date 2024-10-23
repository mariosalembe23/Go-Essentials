package main

import "fmt"

// Função que imprime a string apontada por str
func ft_putstr(str *string) {
	if str == nil {
		fmt.Printf("\n") 
		return
	}
	fmt.Printf("%s", *str) 
}
