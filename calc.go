package main

import "fmt"

func main() {
	x := soma(10, 50, 15)
	y := multiplica(5, 20)
	w := subtrai(88, 100)
	z := divide(10)
	fmt.Println(x, y, w, z)
}
func soma(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}
func subtrai(i ...int) int {
	total := 0
	for _, v := range i {
		total = v - total
	}
	return total
}
func multiplica(i ...int) int {
	total := 1
	for _, v := range i {
		total = total * v
	}
	return total
}
func divide(i ...int) int {
	total := 2
	for _, v := range i {
		total = v / total
	}
	return total
}
