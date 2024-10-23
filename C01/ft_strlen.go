package main

import "fmt"

func ft_strlen(str *string) int {
	new_str := *str
	return len(new_str)
}
