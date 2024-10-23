package main

import  "fmt"

func ft_div_mod(a int, b int, div *int, mod *int) {

	if b == 0 {
		fmt.Println("Erro: Divisão por zero")
		return
	}
	*div = a / b
	*mod = a % b
}