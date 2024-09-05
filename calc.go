package main

import "fmt"

func main() {
	x := somar(10, 50, 15)
	y := multiplicar(5, 20)
	w := subtrair(88, 100)
	z := divisao(10)
	fmt.Println(x, y, w, z)
}
func somar(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}
func subtrair(i ...int) int {
	total := 0
	for _, v := range i {
		total = v - total
	}
	return total
}
func multiplicar(i ...int) int {
	total := 1
	for _, v := range i {
		total = total * v
	}
	return total
}
func divisao(i ...int) int {
	total := 2
	for _, v := range i {
		total = v / total
	}
	return total
}
