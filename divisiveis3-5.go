package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, " - Numero Divisível por 3 e 5")
		} else if i%3 == 0 {
			fmt.Println(i, " - Numero Divisível apenas por 3")
		} else if i%5 == 0 {
			fmt.Println(i, " - Numero Divisível apenas por 5")
		} else {
			fmt.Println(i)
		}
	}
}
